// package main

// import (
// 	"fmt"
// 	"sync"

// )

// func main() {
// 	wg:= sync.WaitGroup{}
// 	counter := 3
// 	for i := 0; i <= counter; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()		
// 			fmt.Println(n * n)
// 		}(i)
	
// 	}
// 	wg.Wait()
// }