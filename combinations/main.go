package main

import "fmt"

func main() {
	fmt.Println(combine(4, 3))
}

func combine(n int, k int) [][]int {
	r := [][]int{}
	putNum(&r, []int{}, 0, n, k)
	return r
}

func putNum(r *[][]int, c []int, start, n, k int) {
	if k == 0 {
		item := make([]int, len(c))
		copy(item, c)
		*r = append(*r, item)
		return
	}
	for i := start; i < n; i++ {
		c = append(c, i+1)
		putNum(r, c, i+1, n, k-1)
		c = c[:len(c)-1]
	}

}

// 1 2 3 4
func combine2(n int, k int) [][]int {
	nums := make([]int, 0)
	r := [][]int{}
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	putNum2(&r, []int{}, nums, k)
	return r
}

func putNum2(r *[][]int, current []int, candidate []int, k int) {
	if k == 0 {
		*r = append(*r, current)
		return
	}

	for index, v := range candidate {
		newCurrent := append([]int{}, current...)
		newCurrent = append(newCurrent, v)
		// newCand := append([]int{}, candidate[:index]...)
		newCand := append([]int{}, candidate[index+1:]...)

		putNum2(r, newCurrent, newCand, k-1)
	}
}
