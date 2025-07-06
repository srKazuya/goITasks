// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func generator(ctx context.Context, numbers []int) chan int {
// 	outCh := make(chan int)

// 	go func() {
// 		defer close(outCh)
// 		for _, num := range numbers {
// 			select {
// 			case <-ctx.Done():
// 			case outCh <- num:
// 			}
// 		}
// 	}()

// 	return outCh
// }
// func add(ctx context.Context, inputCh chan int) chan int {
// 	outCh := make(chan int)
// 	go func() {
// 		defer close(outCh)
// 		for v := range inputCh {
// 			select {
// 			case <-ctx.Done():
// 			case outCh <- v * 2:
// 			}
// 		}
// 	}()
// 	return outCh
// }

// func multiply(ctx context.Context, inChan chan int) chan int {
// 	outCh := make(chan int)
// 	go func() {

// 		defer close(outCh)
// 		for v := range inChan {
// 			select {
// 			case <-ctx.Done():
// 			case outCh <- v + 2:
// 			}
// 		}
// 	}()
// 	return outCh
// }

// func fanOut(ctx context.Context, inputCh chan int, workers int) []chan int {
// 	outCh := make([]chan int, workers)
// 	for i := 0; i < workers; i++ {
// 		outCh[i] = add(ctx, inputCh)

// 	}

// 	return outCh
// }

// func fanIn(ctx context.Context, chans ...chan int) chan int {
// 	outCh := make(chan int)
// 	var wg sync.WaitGroup
// 	for _, ch := range chans {
		
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()
// 			for val := range ch {
// 				select {
// 				case <-ctx.Done():
// 				case outCh <- val:
// 				}
// 			}
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(outCh)
// 	}()

// 	return outCh
// }
// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	inputCh := generator(ctx, numbers)

// 	addChans := fanOut(ctx, inputCh, 10)

// 	addRes := fanIn(ctx, addChans...)

// 	result := add(ctx, addRes)
// 	for v := range result {
// 		fmt.Println(v)
// 	}

// }
