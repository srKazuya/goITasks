// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// 	"runtime"
// )

// func writer() <-chan int {
// 	ch := make(chan int)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		for i := range 5 {
// 			ch <- i + 1
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for i := range 5 {
// 			ch <- i + 11
// 		}
// 	}()

// 	go func ()  {
// 		wg.Wait()
// 		close(ch)
// 	}()
// 	return ch
// }

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	ch := writer()

// 	for {
// 		v, ok := <-ch
// 		if !ok {
// 			break
// 		}
// 		fmt.Println("v = ", v)
// 	}
// 	time.Sleep(1 * time.Second)
// }
