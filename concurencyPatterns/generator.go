// // https://habr.com/ru/articles/852556/
// package main

// import "fmt"

// func main() {
// 	items := []int{20, 30, 40, 50}

// 	dataCh := generator(items)

// 	process(dataCh)
// }

// func generator(items []int) <-chan int {
// 	ch := make(chan int)

// 	go func() {

// 		for _, v := range items {
// 			ch <- v
// 		}

// 		close(ch)
// 	}()
// 	return ch
// }

// func process(ch <-chan int) {
// 	for {
// 		item, ok := <-ch
// 		if !ok {
// 			fmt.Println("Err")
// 			break
// 		}
// 		fmt.Println(item)
// 	}
// }
