package main

import "fmt"

func main() {
	// 1. 배열: 크기가 고정 (타입의 일부)
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr, len(arr))

	// 2. 슬라이스: 크기가 가변적, 실무에서는 배열보다 슬라이스를 훨씬 많이 사용
	s := []int{1, 2, 3}
	s = append(s, 4, 5) // 요소 추가, 필요시 내부 배열 재할당
	fmt.Println(s, len(s), cap(s))

	// 3. make로 슬라이스 생성 (길이, 용량 지정 가능)
	s2 := make([]int, 3, 5) // len=3, cap=5
	fmt.Println(s2, len(s2), cap(s2))

	// 4. 슬라이싱 (slice[start:end], end는 포함 안 됨)
	sub := s[1:3]
	fmt.Println(sub)

	// 5. 2차원 슬라이스
	grid := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(grid[1][0])

	// 6. 맵: key-value 자료구조
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	m["c"] = 3 // 추가/수정

	// 7. 맵 조회 시 comma-ok 패턴으로 존재 여부 확인
	val, ok := m["z"]
	if !ok {
		fmt.Println("z 없음, val:", val) // 없으면 zero value
	}

	// 8. 맵 삭제
	delete(m, "a")

	// 9. 맵 순회 (순서 보장 안 됨!)
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 10. 슬라이스는 참조 타입에 가까움 -> 함수에 넘기면 원본 데이터 공유
	modify(s)
	fmt.Println("after modify:", s)
}

func modify(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}
