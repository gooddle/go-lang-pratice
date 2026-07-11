package main

import "fmt"

func main() {
	// 1. if / else (조건에 괄호 없음, 중괄호는 필수)
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 2. if에 초기화문 포함 (그 안에서만 유효한 변수)
	if v := score * 2; v > 100 {
		fmt.Println("over 100:", v)
	}

	// 3. for는 Go의 유일한 반복문 (while, do-while 없음)
	for i := 0; i < 3; i++ {
		fmt.Println("for:", i)
	}

	// 4. while처럼 사용 (조건만)
	n := 0
	for n < 3 {
		n++
	}

	// 5. 무한 루프 + break
	count := 0
	for {
		count++
		if count >= 3 {
			break
		}
	}

	// 6. range로 슬라이스/맵/문자열 순회
	nums := []int{10, 20, 30}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}

	// 7. switch (fallthrough 없으면 자동 break)
	switch day := 3; day {
	case 1, 7:
		fmt.Println("주말")
	case 2, 3, 4, 5, 6:
		fmt.Println("평일")
	default:
		fmt.Println("알 수 없음")
	}

	// 8. 조건식 switch (switch true와 동일한 효과)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	// 9. continue로 다음 반복으로 스킵
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd:", i)
	}
}
