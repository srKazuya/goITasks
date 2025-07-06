package main

import (
	"fmt"
)

func search(nums []int) []int {
	m := len(nums) / 2
	right := nums[m:]
	return right

}
func main() {
	nums := []int{1, 2, 3, 4, 6, 7}
	fmt.Println(search(nums))
}
