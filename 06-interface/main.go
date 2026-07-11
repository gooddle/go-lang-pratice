package main

import "fmt"

// 1. 인터페이스 선언: 메서드 시그니처 목록
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

//  2. 구조체가 인터페이스를 구현할 때 implements 키워드 없음
//     -> Shape이 요구하는 메서드를 전부 가지고 있으면 자동으로 Shape 타입으로 취급됨 (덕 타이핑)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 3. 인터페이스 타입을 파라미터로 받으면 다형성 구현 가능
func describe(s Shape) {
	fmt.Printf("넓이: %.2f, 둘레: %.2f\n", s.Area(), s.Perimeter())
}

// 4. 빈 인터페이스 (interface{} 또는 any): 어떤 타입이든 담을 수 있음
func printAny(v any) {
	fmt.Println(v)
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
	}

	for _, s := range shapes {
		describe(s)
	}

	// 5. 타입 단언 (type assertion): 인터페이스 값의 실제 타입 꺼내기
	var s Shape = Rectangle{Width: 2, Height: 2}
	if rect, ok := s.(Rectangle); ok {
		fmt.Println("실제 타입은 Rectangle:", rect)
	}

	// 6. type switch: 여러 타입을 분기 처리
	check(shapes[0])
	check(shapes[1])

	printAny(123)
	printAny("hello")
}

func check(s Shape) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Println("사각형, 너비:", v.Width)
	case Circle:
		fmt.Println("원, 반지름:", v.Radius)
	default:
		fmt.Println("알 수 없는 도형")
	}
}
