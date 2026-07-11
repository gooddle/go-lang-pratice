package main

import (
	"fmt"
	"sync"
)

// 1. goroutine: go 키워드만 붙이면 별도의 가벼운 스레드에서 함수 실행
func sayHello(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 3. WaitGroup: 여러 goroutine이 끝날 때까지 기다리는 카운터
	fmt.Println("hello from goroutine", id)
}

func main() {
	// 2. 여러 goroutine 동시 실행 + WaitGroup으로 종료 대기
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sayHello(i, &wg)
	}
	wg.Wait() // 모든 goroutine이 끝날 때까지 블록

	// 4. channel: goroutine 간 데이터를 주고받는 파이프 (make(chan 타입))
	ch := make(chan string)

	go func() {
		ch <- "채널로 보낸 메시지" // 채널에 값 전송
	}()

	msg := <-ch // 채널에서 값 수신 (받을 때까지 블록)
	fmt.Println(msg)

	// 5. 버퍼 채널: 지정한 크기만큼은 블록 없이 전송 가능
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2
	fmt.Println(<-buffered, <-buffered)

	// 6. 채널 닫기 + range로 순회 (닫히면 루프 자동 종료)
	results := make(chan int)
	go func() {
		defer close(results)
		for i := 0; i < 3; i++ {
			results <- i * i
		}
	}()
	for r := range results {
		fmt.Println("결과:", r)
	}

	// 7. select: 여러 채널 중 준비된 것을 처리 (switch의 채널 버전)
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c1 <- "c1에서 옴"

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	default:
		fmt.Println("아무 채널도 준비 안 됨")
	}

	// 8. Mutex: 여러 goroutine이 공유 데이터에 동시 접근할 때 race condition 방지
	var mu sync.Mutex
	counter := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Println("counter:", counter)
}
