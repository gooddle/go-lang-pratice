// chi(경량 라우터) + Bun(경량 ORM) + SQLite 로 만든 Todo REST API 예제.
//
// 설정은 모두 환경변수로 주입한다 (시크릿/DB 경로를 코드에 하드코딩하지 않는다).
//   PORT       : 서버 포트           (기본 8080)
//   DB_DSN     : SQLite DSN/파일 경로 (기본 file:todo.db?cache=shared)
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

// Todo 는 bun 모델. 구조체 태그로 테이블/컬럼을 매핑한다.
type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	Title     string    `bun:"title,notnull" json:"title"`
	Done      bool      `bun:"done,notnull" json:"done"`
	CreatedAt time.Time `bun:"created_at,notnull" json:"created_at"`
}

// server 는 핸들러가 공유하는 의존성(여기서는 DB)을 담는다.
type server struct {
	db *bun.DB
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	port := getenv("PORT", "8080")
	dsn := getenv("DB_DSN", "file:todo.db?cache=shared")

	// sqliteshim 은 CGO 없이(순수 Go) 동작하는 SQLite 드라이버를 물어온다.
	sqldb, err := sql.Open(sqliteshim.ShimName, dsn)
	if err != nil {
		return err
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, sqlitedialect.New())

	ctx := context.Background()
	if _, err := db.NewCreateTable().Model((*Todo)(nil)).IfNotExists().Exec(ctx); err != nil {
		return err
	}

	srv := &server{db: db}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", srv.listTodos)
		r.Post("/", srv.createTodo)
		r.Get("/{id}", srv.getTodo)
		r.Put("/{id}", srv.updateTodo)
		r.Delete("/{id}", srv.deleteTodo)
	})

	log.Printf("listening on :%s", port)
	return http.ListenAndServe(":"+port, r)
}

func (s *server) listTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	if err := s.db.NewSelect().Model(&todos).Order("id ASC").Scan(r.Context()); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusOK, todos)
}

func (s *server) createTodo(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if in.Title == "" {
		writeError(w, http.StatusBadRequest, errors.New("title is required"))
		return
	}

	todo := &Todo{Title: in.Title, CreatedAt: time.Now()}
	if _, err := s.db.NewInsert().Model(todo).Exec(r.Context()); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusCreated, todo)
}

func (s *server) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := idParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	todo := new(Todo)
	err = s.db.NewSelect().Model(todo).Where("id = ?", id).Scan(r.Context())
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, errors.New("todo not found"))
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

func (s *server) updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := idParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var in struct {
		Title *string `json:"title"`
		Done  *bool   `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	q := s.db.NewUpdate().Model((*Todo)(nil)).Where("id = ?", id)
	if in.Title != nil {
		q = q.Set("title = ?", *in.Title)
	}
	if in.Done != nil {
		q = q.Set("done = ?", *in.Done)
	}

	res, err := q.Exec(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if n, _ := res.RowsAffected(); n == 0 {
		writeError(w, http.StatusNotFound, errors.New("todo not found"))
		return
	}

	todo := new(Todo)
	if err := s.db.NewSelect().Model(todo).Where("id = ?", id).Scan(r.Context()); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

func (s *server) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := idParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	res, err := s.db.NewDelete().Model((*Todo)(nil)).Where("id = ?", id).Exec(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if n, _ := res.RowsAffected(); n == 0 {
		writeError(w, http.StatusNotFound, errors.New("todo not found"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- helpers ---

func idParam(r *http.Request) (int64, error) {
	return strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}
