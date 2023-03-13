package main

import "fmt"

func main() {
	println(getSkyline([][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}))
	// println(getSkyline([][]int{
	// 	{2, 5, 10},
	// 	{5, 7, 10},
	// }))
}

func getSkyline(buildings [][]int) [][]int {
	// ❌
	// nodes := [][]int{}
	// for _, v := range buildings {
	// 	nodes = append(nodes, []int{v[0], v[2]})
	// 	nodes = append(nodes, []int{v[1], 0})
	// }
	return getSkylineImpl(buildings)
}

func getSkylineImpl(buildings [][]int) [][]int {
	if len(buildings) < 2 {
		return [][]int{
			{buildings[0][0], buildings[0][2]},
			{buildings[0][1], 0},
		}
	}
	mid := len(buildings) / 2
	left := getSkylineImpl(buildings[:mid])
	right := getSkylineImpl(buildings[mid:])

	i, j := 0, 0
	result := [][]int{}
	preH1 := 0
	preH2 := 0
	x := 0
	pre := 0
	fmt.Println(left, right)
	for i < len(left) && j < len(right) {
		p1  := left[i]
		p2 := right[j]
		if p1[0] < p2[0] {
			x = p1[0]
			preH1 = p1[1]
			i++
		} else if p1[0] > p2[0] {
			x = p2[0]
			preH2 = p2[1]
			j++
		} else {
			x = p1[0]
			preH1 = p1[1]
			preH2 = p2[1]
			i++
			j++
		}
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		maxH := max(preH1, preH2)
		if maxH != pre {
			result = append(result, []int{x, maxH})
			pre = maxH

		}


	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
