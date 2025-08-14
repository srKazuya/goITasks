// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// type Taxi struct {
// 	pos  int64
// 	time int64
// 	seen bool
// }

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	var quantity int
// 	var L, S int64
// 	if _, err := fmt.Fscan(in, &quantity, &L, &S); err != nil {
// 		return
// 	}

// 	taxis := make(map[int]Taxi)

// 	for i := 0; i < quantity; i++ {
// 		var cmd string
// 		fmt.Fscan(in, &cmd)
// 		switch cmd {
// 		case "TAXI":
// 			taxiCommand(in, taxis)
// 		case "ORDER":
// 			orderCommand(in, out, taxis, L, S)
// 		}
// 	}
// }

// func taxiCommand(in *bufio.Reader, taxis map[int]Taxi) {
// 	var ts int64
// 	var id int
// 	var pos int64
// 	fmt.Fscan(in, &ts, &id, &pos)
// 	taxis[id] = Taxi{pos: pos, time: ts, seen: true}
// }

// func orderCommand(in *bufio.Reader, out *bufio.Writer, taxis map[int]Taxi, L, S int64) {
// 	var ts, orderPos, orderTime int64
// 	var orderID int
// 	fmt.Fscan(in, &ts, &orderID, &orderPos, &orderTime)

// 	res := []int{}
// 	for id, t := range taxis {
// 		if !t.seen || t.time > ts {
// 			continue
// 		}

// 		dt := ts - t.time
// 		if dt < 0 {
// 			continue
// 		}

// 		maxTravel := dt * S
// 		if maxTravel > L {
// 			maxTravel = L
// 		}
// 		lenArc := maxTravel

// 		start := t.pos % L
// 		var y0 int64 = (orderPos - start + L) % L

// 		var worstDist int64
// 		if lenArc >= y0+1 { 
// 			worstDist = L - 1
// 		} else {
// 			worstDist = y0
// 		}

// 		worstTime := (worstDist + S - 1) / S
// 		if worstTime <= orderTime {
// 			res = append(res, id)
// 			if len(res) == 5 {
// 				break
// 			}
// 		}
// 	}

// 	if len(res) == 0 {
// 		fmt.Fprintln(out, -1)
// 		return
// 	}

// 	sort.Ints(res)
// 	for i, v := range res {
// 		if i > 0 {
// 			fmt.Fprint(out, " ")
// 		}
// 		fmt.Fprint(out, v)
// 	}
// 	fmt.Fprintln(out)
// }
