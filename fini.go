// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
// )

// type Point struct{ x, y int }

// func main() {
//     in := bufio.NewReader(os.Stdin)
//     var n, m, d int
//     fmt.Fscan(in, &n, &m, &d)

//     grid := make([][]byte, n)
//     for i := range grid {
//         var s string
//         fmt.Fscan(in, &s)
//         grid[i] = []byte(s)
//     }

//     dist := make([][]int, n)
//     for i := range dist {
//         dist[i] = make([]int, m)
//         for j := range dist[i] {
//             dist[i][j] = -1
//         }
//     }

//     q := make([]Point, 0, n*m)
//     for i := 0; i < n; i++ {
//         for j := 0; j < m; j++ {
//             if grid[i][j] == 'x' || grid[i][j] == 'X' {
//                 dist[i][j] = 0
//                 q = append(q, Point{i, j})
//             }
//         }
//     }

//     dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
//     for head := 0; head < len(q); head++ {
//         p := q[head]
//         for _, dir := range dirs {
//             nx, ny := p.x+dir.x, p.y+dir.y
//             if nx >= 0 && nx < n && ny >= 0 && ny < m && dist[nx][ny] == -1 {
//                 dist[nx][ny] = dist[p.x][p.y] + 1
//                 q = append(q, Point{nx, ny})
//             }
//         }
//     }

//     dp := make([][]int, n)
//     maxSide := 0
//     for i := range dp {
//         dp[i] = make([]int, m)
//     }
//     for i := 0; i < n; i++ {
//         for j := 0; j < m; j++ {
//             if dist[i][j] >= d {
//                 if i == 0 || j == 0 {
//                     dp[i][j] = 1
//                 } else {
//                     dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
//                 }
//                 if dp[i][j] > maxSide {
//                     maxSide = dp[i][j]
//                 }
//             }
//         }
//     }

//     fmt.Println(maxSide)
// }

// func min3(a, b, c int) int {
//     if a < b {
//         if a < c {
//             return a
//         }
//         return c
//     }
//     if b < c {
//         return b
//     }
//     return c
// }
