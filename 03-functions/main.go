package main

import "fmt"

// 1. 기본 함수 선언: func 이름(파라미터) 리턴타입
func add(a int, b int) int {
	return a + b
}

// 2. 같은 타입 파라미터는 마지막에 한번만 타입 명시 가능
func multiply(a, b int) int {
	return a * b
}

// 3. 여러 값 리턴 (Go의 대표적 특징)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("0으로 나눌 수 없음")
	}
	return a / b, nil
}

// 4. named return (리턴 변수 이름 미리 선언, return만 써도 됨)
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // 명시적 값 없이 return -> min, max 반환
}

// 5. 가변 인자 (variadic)
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 6. 함수를 값으로 사용 (고차 함수)
func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

// 7. 클로저 (함수가 외부 변수를 캡처)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(multiply(3, 4))

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("에러:", err)
	} else {
		fmt.Println("결과:", result)
	}

	min, max := minMax([]int{5, 3, 9, 1})
	fmt.Println(min, max)

	fmt.Println(sum(1, 2, 3, 4))

	double := func(x int) int { return x * 2 } // 익명 함수
	fmt.Println(applyTwice(double, 3))

	next := counter()
	fmt.Println(next(), next(), next()) // 1 2 3
}
