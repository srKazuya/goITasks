// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func randomTimeWork() {
// 	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
// }

// func predicatableTimeWork() {
// 	ch := make(chan struct{})

// 	go func() {
// 		randomTimeWork()
// 		close(ch)
// 	}()

// 	timer := time.NewTimer(3 * time.Second)
// 	defer timer.Stop()
// 	select {
// 	case <-timer.C:
// 		fmt.Println("Ошибк")
// 	case <-ch:
// 		fmt.Println("ok")
// 	}
// }

// func main() {
// 	predicatableTimeWork()
// }
