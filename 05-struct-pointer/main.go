package main

import "fmt"

// 1. 구조체 정의
type Person struct {
	Name string
	Age  int
}

// 2. 값 리시버 메서드: 구조체를 복사해서 사용, 원본 변경 안 됨
func (p Person) Greet() string {
	return fmt.Sprintf("안녕, 나는 %s, %d살이야", p.Name, p.Age)
}

// 3. 포인터 리시버 메서드: 원본을 직접 수정 가능
func (p *Person) Birthday() {
	p.Age++
}

// 4. 구조체 임베딩 (상속 대신 조합)
type Employee struct {
	Person  // 임베딩: Person의 필드/메서드를 그대로 사용 가능
	Company string
}

func main() {
	// 5. 구조체 생성
	p1 := Person{Name: "철수", Age: 20}
	p2 := Person{"영희", 22} // 필드명 생략 (순서대로)

	fmt.Println(p1.Greet())
	p1.Birthday() // Go가 자동으로 &p1 처리
	fmt.Println(p1.Age)

	// 6. 포인터: 변수의 주소를 저장하는 타입
	var ptr *Person = &p2
	ptr.Age = 30 // 포인터를 통해 원본 수정 (*ptr).Age와 동일하게 동작
	fmt.Println(p2.Age)

	// 7. new()로 포인터 생성 (제로값으로 초기화된 구조체의 포인터)
	p3 := new(Person)
	p3.Name = "민수"

	// 8. 임베딩된 구조체 사용
	e := Employee{
		Person:  Person{Name: "지훈", Age: 28},
		Company: "구들컴퍼니",
	}
	fmt.Println(e.Greet()) // Person의 메서드를 바로 호출 가능
	fmt.Println(e.Name, e.Company)

	fmt.Println(p3)
}
