package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {

	result := [][]int{}
	placeNum(&result, []int{}, nums)
	return result
}

func placeNum(r *[][]int, current []int, rem []int) {

	item := make([]int, len(current))
	copy(item, current)
	*r = append(*r, item)
	if len(rem) == 0 {
		return
	}
	for index, num := range rem {
		current = append(current, num)
		placeNum(r, current, rem[index+1:])
		current = current[:len(current)-1]
	}

}
