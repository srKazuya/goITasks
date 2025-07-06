// package main

// import "fmt"

// func generator(doneCh chan struct{}, nums []int) chan int {
// 	outCh := make(chan int)
// 	go func() {
// 		for _, v := range nums {
// 			select {
// 			case <-doneCh:
// 				return
// 			case outCh <- v:
// 			}
// 		}
// 		close(outCh)
// 	}()

// 	return outCh
// }

// func add(doneCh chan struct{}, inCh chan int) chan int {
// 	outCh := make(chan int)
// 	go func() {
// 		defer close(outCh)
// 		for val := range inCh {
// 			select {
// 			case <-doneCh:
// 				return
// 			case outCh <- val + 2:
// 			}
// 		}
// 	}()
// 	return outCh
// }

// func multiply(doneCh chan struct{}, inCh chan int) chan int {
// 	ch := make(chan int)
// 	go func() {
// 		defer close(ch)
// 		for val := range inCh {
// 			select {
// 			case <-doneCh:
// 				return
// 			case ch <- val * 2:
// 			}
// 		}
// 	}()
// 	return ch
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}

// 	doneCh := make(chan struct{})
// 	defer close(doneCh)

// 	inputCh := generator(doneCh, numbers)

// 	addCh := add(doneCh, inputCh)

// 	multCh := multiply(doneCh, addCh)
// 	for v := range multCh {

// 		fmt.Println(v)
// 	}

// }
