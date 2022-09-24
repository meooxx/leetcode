package main

import (
	"fmt"
)

//import "math"

// func twoSum(nums []int, target int) []int {
// 	// n := len(nums)
// 	m := make(map[int]int)
// 	//var m map[int] int
// 	for i := 0; i < len(nums); i++ {
// 		//v:=nums[i]

// 		rest := target - nums[i]
// 		if _, ok := m[rest]; ok {
// 			return []int{m[rest], i}
// 		}

// 		m[nums[i]] = i
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	s := map[int]int{}
	r := []int{}
	for index, v := range nums {
		surplus := target - v
		if resultIndex, ok := s[v]; ok {
			r = append(r, index, resultIndex)
			return r
		} else {
			s[surplus] = index
		}
	}
	return []int{}
}
func main() {
	s := twoSum([]int{-1, 2, 1, 3}, 4)
	fmt.Println(s)

}
