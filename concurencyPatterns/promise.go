
// package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// type Result struct {
// 	val int
// 	err error
// }

// func Promise(task func() (int, error)) <-chan Result {

// 	ch := make(chan Result)

// 	go func() {
// 		val, err := task()
// 		ch <- Result{val: val, err: err}
// 		close(ch)
// 	}()

// 	return ch
// }

// func main() {

// 	longRunningTask := func() (int, error) {
// 		time.Sleep(1 * time.Second)
// 		return 42, errors.New("Ошибка работы")
// 	}

// 	future := Promise(longRunningTask)

// 	fmt.Println("Задача запушена...")

// 	for {
// 		result, ok := <-future
// 		if !ok {
// 			break
// 		}
// 		if result.err != nil {
// 			fmt.Println("Ошибка:", result.err)
// 		} else {
// 			fmt.Println("Результат:", result.val)
// 		}
// 	}
// }
