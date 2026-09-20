# 09 - chi + Bun + SQLite

경량 라우터 [chi](https://github.com/go-chi/chi)와 경량 ORM [Bun](https://bun.uptrace.dev/)을
SQLite와 함께 사용한 Todo REST API 예제입니다.

## 스택

- **chi** — 표준 `net/http`에 얹는 가벼운 라우터/미들웨어
- **Bun** — SQL에 가까우면서 구조체 매핑을 제공하는 경량 ORM
- **SQLite** — `sqliteshim` 드라이버로 CGO 없이(순수 Go) 동작

## 설정 (환경변수)

시크릿/경로는 코드에 하드코딩하지 않고 환경변수로 주입합니다. (`.env` 및 `*.db` 는 커밋하지 않습니다 — `.gitignore` 처리됨)

| 변수 | 설명 | 기본값 |
|------|------|--------|
| `PORT` | 서버 포트 | `8080` |
| `DB_DSN` | SQLite DSN/파일 경로 | `file:todo.db?cache=shared` |

## 실행

```bash
# 모듈 루트에서
go run ./09-chi-bun-sqlite

# 또는 환경변수 지정
PORT=9000 DB_DSN="file:mydata.db" go run ./09-chi-bun-sqlite
```

## API

| 메서드 | 경로 | 설명 |
|--------|------|------|
| GET | `/todos` | 목록 조회 |
| POST | `/todos` | 생성 (`{"title": "..."}`) |
| GET | `/todos/{id}` | 단건 조회 |
| PUT | `/todos/{id}` | 수정 (`{"title": "...", "done": true}`) |
| DELETE | `/todos/{id}` | 삭제 |

### 예시

```bash
curl -X POST localhost:8080/todos -d '{"title":"go 공부"}'
curl localhost:8080/todos
curl -X PUT localhost:8080/todos/1 -d '{"done":true}'
curl -X DELETE localhost:8080/todos/1
```
