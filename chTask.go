package main

import (
	"fmt"
	"time"
)

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func double(writerCh <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for v := range writerCh {
			ch <- v * 2
		}
		close(ch)
	}()
	return ch
}

func reader(doubleCh <-chan int) {
	for i := range doubleCh {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	reader(double(writer()))
}