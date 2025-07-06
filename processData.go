// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func processData(ctx context.Context, val int) int {
// 	ch := make(chan struct{})
// 	go func() {

// 		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
// 		close(ch)
// 	}()
// 	select {
// 	case <-ch:

// 	}
// 	return val * 2
// }

// func main() {
// 	in := make(chan int)
// 	out := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			in <- i
// 		}
// 		close(in)
// 	}()

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	now := time.Now()
// 	processParallel(ctx, in, out, 5)

// 	for val := range out {
// 		fmt.Println(val)
// 	}

// 	fmt.Println(time.Since(now))
// }

// func processParallel(ctx context.Context, in <-chan int, out chan<- int, numWorkers int) {
// 	wg := &sync.WaitGroup{}
// 	for i := 0; i < numWorkers; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for i := range in {
// 				select {
// 				case out <- processData(ctx, i):
// 				case <-ctx.Done():
// 					return
// 				}

// 			}
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// }

// // package main

// // import (
// // 	"fmt"
// // 	"math/rand"
// // 	"time"
// // )

// // func processData(val int) int {
// // 	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
// // 	return val * 2
// // }

// // func main() {
// // 	in := make(chan int)
// // 	out := make(chan int)

// // 	go func() {
// // 		for i := range 100 {
// // 			in <- i
// // 		}
// // 		close(in)
// // 	}()

// // 	now := time.Now()
// // 	processParallel(in, out, 5)

// // 	for val := range out {
// // 		fmt.Println(val)
// // 	}

// // 	fmt.Println(time.Since(now))
// // }

// // func processParallel(in <-chan int, out chan <- int, numWorkers int) {

// // }
