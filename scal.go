// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	var vectorLen int
// 	fmt.Fscan(in, &vectorLen)

// 	Q := make([]int64, vectorLen)
// 	C := make([]int64, vectorLen)
// 	for i := 0; i < vectorLen; i++ {
// 		fmt.Fscan(in, &Q[i])
// 	}
// 	for i := 0; i < vectorLen; i++ {
// 		fmt.Fscan(in, &C[i])
// 	}

// 	var A, B int64
// 	fmt.Fscan(in, &A, &B)

// 	var res int64

// 	if A < B {
// 		for i := 0; i < vectorLen; i++ {

// 			D := ((C[i] * (B - A)) / 255) + A
// 			res += Q[i] * D
// 		}
// 	} else {
// 		for i := 0; i < vectorLen; i++{
// 			res += Q[i] * 0
// 		}
// 	}

// 	fmt.Fprintln(out, res)
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func ceilDiv(a, b int64) int64 {
// 	if a >= 0 {
// 		return (a + b - 1) / b
// 	}
// 	return a / b 
// }

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	var vectorLen int
// 	fmt.Fscan(in, &vectorLen)

// 	Q := make([]int64, vectorLen)
// 	C := make([]int64, vectorLen)
// 	for i := 0; i < vectorLen; i++ {
// 		fmt.Fscan(in, &Q[i])
// 	}
// 	for i := 0; i < vectorLen; i++ {
// 		fmt.Fscan(in, &C[i])
// 	}

// 	var A, B int64
// 	fmt.Fscan(in, &A, &B)

// 	var res int64
// 	diff := B - A

// 	if A < B {
// 		for i := 0; i < vectorLen; i++ {
// 			ci := C[i]
// 			l := A + ceilDiv(ci*diff, 255)
// 			num := (ci+1)*diff - 1
// 			h := A + num/255
// 			if h < l {
// 				h = l
// 			}
// 			D := (l + h) / 2 
// 			res += Q[i] * D
// 		}
// 	} else {
		
// 		for i := 0; i < vectorLen; i++ {
// 			res += Q[i] * A
// 		}
// 	}

// 	fmt.Fprintln(out, res)
// }
