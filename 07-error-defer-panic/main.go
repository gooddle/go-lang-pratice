package main

import (
	"errors"
	"fmt"
)

// 1. Go는 예외(exception)가 없음. 에러는 그냥 값(error 타입)이고 반환값으로 전달함
var ErrNotFound = errors.New("항목을 찾을 수 없음") // 2. 커스텀 에러 값 정의

func findUser(id int) (string, error) {
	if id != 1 {
		return "", ErrNotFound
	}
	return "철수", nil
}

// 3. 커스텀 에러 타입 (Error() 메서드만 구현하면 error 인터페이스를 만족함)
type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s 필드가 유효하지 않음", e.Field)
}

func validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age"}
	}
	return nil
}

//  4. defer: 함수가 끝날 때(return 직전) 실행됨. 리소스 정리(파일 닫기 등)에 주로 사용
//  5. panic/recover: 프로그램의 심각한 오류 처리. 일반적인 에러 처리에는 error를 쓰고,
//     panic은 복구 불가능한 상황에서만 제한적으로 사용
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil { // panic을 잡아서 error로 변환
			err = fmt.Errorf("복구됨: %v", r)
		}
	}()

	result = a / b // b가 0이면 런타임 panic 발생
	return result, nil
}

func main() {
	// 6. 에러 체크는 항상 if err != nil 패턴
	_, err := findUser(2)
	if err != nil {
		// 7. errors.Is로 특정 에러 값인지 비교
		if errors.Is(err, ErrNotFound) {
			fmt.Println("사용자 없음")
		}
	}

	if err := validate(-1); err != nil {
		// 8. errors.As로 특정 에러 타입으로 꺼내기
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("검증 실패:", ve.Field)
		}
	}

	defer fmt.Println("1. 가장 나중에 출력됨 (defer는 LIFO)")
	defer fmt.Println("2. 그 다음 나중에 출력됨")
	fmt.Println("3. 먼저 출력됨")

	res, err := safeDivide(10, 0)
	fmt.Println(res, err)
}
