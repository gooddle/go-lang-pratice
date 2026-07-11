# Go 기본 문법 정리

Go(Golang) 기본 문법을 주제별 예제 코드로 정리한 레포입니다.
각 폴더는 `go run ./폴더명` 으로 바로 실행해볼 수 있습니다.

## 목차

| 폴더 | 주제 | 내용 |
|---|---|---|
| [01-variables](01-variables/main.go) | 변수, 상수, 기본 타입 | `var`, `:=`, `const`, `iota`, 제로값 |
| [02-control-flow](02-control-flow/main.go) | 제어문 | `if`, `for`(유일한 반복문), `switch`, `range` |
| [03-functions](03-functions/main.go) | 함수 | 다중 리턴값, named return, 가변 인자, 클로저 |
| [04-arrays-slices-maps](04-arrays-slices-maps/main.go) | 자료구조 | 배열 vs 슬라이스, `append`, `make`, 맵, comma-ok |
| [05-struct-pointer](05-struct-pointer/main.go) | 구조체, 포인터 | 값 리시버/포인터 리시버, 구조체 임베딩 |
| [06-interface](06-interface/main.go) | 인터페이스 | 덕 타이핑, 타입 단언, type switch, `any` |
| [07-error-defer-panic](07-error-defer-panic/main.go) | 에러 처리 | `error` 값, 커스텀 에러, `defer`, `panic`/`recover` |
| [08-goroutine-channel](08-goroutine-channel/main.go) | 동시성 | `goroutine`, `channel`, `select`, `sync.WaitGroup`/`Mutex` |

## Go 문법 핵심 요약

### 1. 변수/상수
- `var name type = value` 또는 타입 추론 `var name = value`
- 함수 내부에서는 `name := value` (짧은 선언)로 대부분 대체 가능
- `const`는 컴파일 타임 상수, `iota`로 연속된 상수 그룹 생성
- 선언만 하고 초기화 안 하면 **제로값**(0, "", false, nil 등)으로 채워짐

### 2. 제어문
- 조건에 괄호 없음, 중괄호는 필수
- 반복문은 `for` 하나뿐 (while, do-while 없음) — 조건만 쓰면 while처럼 동작
- `switch`는 case마다 자동으로 break (fallthrough를 명시해야 다음 case로 이어짐)

### 3. 함수
- 여러 값을 동시에 리턴 가능: `func f() (int, error)`
- 관용적으로 마지막 리턴값을 `error`로 두고 호출부에서 `if err != nil` 체크
- 함수는 1급 시민 — 변수에 담거나 파라미터로 전달 가능 (클로저 지원)

### 4. 슬라이스 / 맵
- 배열은 크기 고정, 슬라이스는 가변 길이 (실무에서는 슬라이스가 기본)
- `append`로 요소 추가, `make([]T, len, cap)`로 생성
- 맵 조회는 `v, ok := m[key]` 패턴으로 존재 여부 확인 (comma-ok)
- 슬라이스/맵은 참조 타입에 가까워서 함수에 넘겨도 원본이 공유됨

### 5. 구조체 / 포인터
- Go에는 클래스가 없고 `struct` + 메서드로 객체지향 유사하게 구현
- 메서드는 리시버로 정의: 값 리시버(복사본), 포인터 리시버(원본 수정)
- 상속 대신 구조체 **임베딩**으로 조합(composition) 사용
- `&`로 주소를 얻고 `*`로 역참조 (C 포인터와 유사하지만 포인터 연산은 없음)

### 6. 인터페이스
- 명시적 `implements` 키워드 없이, 필요한 메서드를 구현하면 자동으로 해당 인터페이스 타입이 됨 (덕 타이핑)
- 빈 인터페이스 `any`(`interface{}`)는 모든 타입을 담을 수 있음
- 타입 단언 `v.(T)`, `type switch`로 실제 타입 분기

### 7. 에러 처리
- 예외(try/catch)가 없음. 에러는 마지막 리턴값으로 전달되는 평범한 값
- `errors.New`, `fmt.Errorf`로 에러 생성, `Error() string` 메서드로 커스텀 에러 타입 정의
- `errors.Is`(값 비교), `errors.As`(타입 추출)
- `defer`는 함수 종료 직전 실행(리소스 정리), `panic`/`recover`는 예외적인 상황에서만 제한적으로 사용

### 8. 동시성 (Go의 핵심 강점)
- `go 함수()` 한 줄로 goroutine(경량 스레드) 실행
- `channel`로 goroutine 간 데이터 주고받기 (`ch <- v`, `v := <-ch`)
- `select`로 여러 채널 중 준비된 것을 처리
- `sync.WaitGroup`으로 goroutine 종료 대기, `sync.Mutex`로 공유 자원 동시 접근 제어

## 실행 방법

```bash
go run ./01-variables
go run ./02-control-flow
# ... 폴더명만 바꿔서 실행
```

## 참고
- Go 공식 튜토리얼: https://go.dev/tour/
- Effective Go: https://go.dev/doc/effective_go
