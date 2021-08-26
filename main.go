package main

import (
	"fmt"
	"time"
)

func send(state chan int) {
	for i := range [10]int{} {
		fmt.Printf(">> sending %d << \n", i)
		state <- i
		fmt.Printf(">> sent %d << \n", i)
	}
	close(state)
	//state <- 0  채널이 닫힌상태에선 보내주면 안된다.
}

// (c chan <- int) 보내기 전용 채널
// (c <- chan int) 받기전용 채널로 변화한다. 이럴땐 함수내에서   <- 쓰면 안된다.
func receive(c <-chan int) {
	for {
		time.Sleep(10 * time.Second)
		a, ok := <-c
		if !ok {
			fmt.Println("end")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

func main() {
	c := make(chan int, 10)
	go send(c)
	fmt.Println("blocking")
	receive(c)

}
