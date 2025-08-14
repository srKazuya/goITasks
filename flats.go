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

// 	var n, m, x, y int
// 	fmt.Fscan(in, &n, &m, &x, &y)

// 	counts := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		counts[i] = make([]int, m)
// 	}

// 	totalLines := n * x
// 	for lineIdx := 0; lineIdx < totalLines; lineIdx++ {
// 		var s string
// 		fmt.Fscan(in, &s) 

// 		floor := lineIdx / x
// 		for apt := 0; apt < m; apt++ {
// 			for k := 0; k < y; k++ {
// 				if s[apt*y+k] == 'X' {
// 					counts[floor][apt]++
// 				}
// 			}
// 		}
// 	}

// 	need := (x*y + 1) / 2
// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if counts[i][j] >= need {
// 				ans++
// 			}
// 		}
// 	}

// 	fmt.Fprintln(out, ans)
// }
