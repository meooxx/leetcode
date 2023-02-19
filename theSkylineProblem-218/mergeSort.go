package main

import "fmt"

func main() {
	fmt.Println(getSkyline([][]int{
		{2, 5, 15},
		{3, 7, 10},
	}))
}

func getSkyline(buildings [][]int) [][]int {
	return getSkylineInner(buildings, 0, len(buildings))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getSkylineInner(buildings [][]int, left, right int) [][]int {
	if right <= left {
		return [][]int{}
	}
	if right-left == 1 {
		return [][]int{
			{buildings[left][0], buildings[left][2]},
			{buildings[left][1], 0},
		}
	}
	mid := (right-left)/2 + left
	leftPoints := getSkylineInner(buildings, left, mid)
	rightPoints := getSkylineInner(buildings, mid, right)

	leftIndex, rightIndex := 0, 0
	preH := 0
	// h1, h2 stand for 
	// 3,15, 6,0, 5,10, 7,0
	// 3,15 vs 5,10  => 3,15(h1=15,h2=10 => 15)
	// 6,0, vs 5,10  => 5,10(h1=0, h2 = 10 => 10)
	h1 := 0
	h2 := 0
	minLeft := 0
	result := [][]int{}
	for leftIndex < len(leftPoints) && rightIndex < len(rightPoints) {
		if leftPoints[leftIndex][0] < rightPoints[rightIndex][0] {
			h1 = leftPoints[leftIndex][1]
			minLeft = leftPoints[leftIndex][0]
			leftIndex++
		} else if leftPoints[leftIndex][0] > rightPoints[rightIndex][0] {
			h2 = rightPoints[rightIndex][1]
			minLeft = rightPoints[rightIndex][0]
			rightIndex++
		} else {
			h2 = rightPoints[rightIndex][1]
			h1 = leftPoints[leftIndex][1]
			minLeft = leftPoints[leftIndex][0]
			leftIndex++
			rightIndex++
		}
		maxH := max(h1, h2)
		if maxH != preH {
			preH = maxH
			result = append(result, []int{
				minLeft, maxH,
			})
		}
	}
	result = append(result, leftPoints[leftIndex:]...)
	result = append(result, rightPoints[rightIndex:]...)
	return result
}
