package main

import "fmt"

func main() {
	// 1. var로 선언 (타입 명시)
	var name string = "gopher"
	var age int = 5

	// 2. 타입 추론 (var, 타입 생략 가능)
	var isActive = true

	// 3. 짧은 선언 (함수 내부에서만 사용 가능, := )
	score := 100

	// 4. 여러 개 한번에 선언
	var (
		width  = 10
		height = 20
	)
	x, y := 1, 2

	// 5. 상수 (const, 실행 중 값 변경 불가)
	const pi = 3.14

	// 6. iota를 이용한 상수 그룹 (0부터 자동 증가)
	const (
		Sunday  = iota // 0
		Monday         // 1
		Tuesday        // 2
	)

	// 7. 기본 타입: bool, string
	//    숫자: int, int8/16/32/64, uint, float32/float64
	var count int64 = 1000000
	var ratio float64 = 3.14

	// 8. 제로값(zero value): 선언만 하면 타입별 기본값으로 초기화됨
	var zeroInt int    // 0
	var zeroStr string // ""
	var zeroBool bool  // false

	fmt.Println(name, age, isActive, score, width, height, x, y, pi)
	fmt.Println(Sunday, Monday, Tuesday, count, ratio)
	fmt.Println(zeroInt, zeroStr, zeroBool)
}
