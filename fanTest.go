// func mergeChan(cs ...<-chan int) <-chan int {
// 	outCh := make(chan int)
// 	var wg sync.WaitGroup
// 	for _, ch := range cs {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for v: =range ch{
// 				outCh <- val
// 			}
// 		}()
// 	}
// 	go func() {
// 			wg.Wait()
// 	close(outCh)
// 	}()
// 	return outCh
// }