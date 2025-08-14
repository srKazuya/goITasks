// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// )

// func main() {
// 	alreadyStored := make(map[int]struct{})
// 	mu := sync.Mutex{}

// 	capacity := 1000

// 	doubles := make([]int, 0, capacity)
// 	for i := 0; i < capacity; i++ {
// 		doubles = append(doubles, rand.Intn(10))
// 	}

// 	uniqueIDs := make(chan int, capacity)
// 	wg := sync.WaitGroup{}


// 	for i := 0; i < capacity; i++ {
// 		i := i
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			if _, ok := alreadyStored[doubles[i]]; !ok {
// 				mu.Lock()
// 				alreadyStored[doubles[i]] = struct{}{}
// 				mu.Unlock()

// 				uniqueIDs <- doubles[i]
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	for val := range uniqueIDs {
// 		fmt.Println(val)
// 	}
// }
