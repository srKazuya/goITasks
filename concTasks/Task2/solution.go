// package main

// import (
// 	"fmt"
// 	"sync"

// )



// func main() {
// 	writes :=1000
// 	var storage map[int]int
// 	storage = make(map[int]int)
	
// 	var mu sync.Mutex
	
// 	wg:= sync.WaitGroup{}
	
// 	wg.Add(writes)
// 	for i:=0; i<writes; i++{
// 		i:=i
		
// 		go func(){
// 			mu.Lock()
// 			defer wg.Done()
// 			storage[i]=i
// 			mu.Unlock()
// 		}()
		
// 	}
// 	wg.Wait()
// 	fmt.Println(storage)
// }
