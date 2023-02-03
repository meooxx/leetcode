package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}}))
	// fmt.Println(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}))

}

func getSkyline1(buildings [][]int) [][]int {
	points := [][]int{}
	for i := range buildings {
		point := buildings[i]
		points = append(points, []int{point[0], -point[2]})
		points = append(points, []int{point[1], point[2]})
	}
	sort.Slice(points, func(i int, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]

	})
	queue := PriorityQueue{head: &PriorityQueue{}}
	result := [][]int{}
	pre := 0
	for i := range points {
		point := points[i]
		if point[1] < 0 {
			queue.Insert(-point[1])
		} else {
			queue.Remove(point[1])
		}
		cur := queue.Peek()
		if cur != pre {
			pre = cur
			result = append(result, []int{point[0], cur})
		}
	}
	return result
}
func getSkyline(buildings [][]int) [][]int {

	// Conquer: If one buiiling is left,
	// return its diagonal points
	if len(buildings) == 1 {
		bldng := buildings[0]
		skyline := [][]int{
			{bldng[0], bldng[2]},
			{bldng[1], 0},
		}
		return skyline
	}

	// Divide:
	mid := len(buildings) / 2
	skyline1 := getSkyline(buildings[:mid])
	skyline2 := getSkyline(buildings[mid:])

	// Combine:
	var h1, h2, prevHeight, x, idx1, idx2 int
	var skyline [][]int

	for idx1 < len(skyline1) && idx2 < len(skyline2) {

		point1 := skyline1[idx1]
		point2 := skyline2[idx2]

		// set x to the min(point1x, point2x)
		// and update previous heights
		switch {
		case point1[0] < point2[0]:
			x, h1 = point1[0], point1[1]
			idx1++
		case point1[0] > point2[0]:
			x, h2 = point2[0], point2[1]
			idx2++
		default:
			x = point1[0]
			h1, h2 = point1[1], point2[1]
			idx1++
			idx2++
		}
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		// currHeight = max between height of
		// the point with the lesser x value
		// and height of the point from the
		// other skyline
		currHeight := max(h1, h2)
		if currHeight != prevHeight {
			skyline = append(skyline, []int{x, currHeight})
		}

		prevHeight = currHeight
	}

	
	// append the rest e.g. 
	// [3,15], [7,12]
	// the rest is               [12, 0]
	// simply append [12, 0] to skyline
	skyline = append(skyline, skyline1[idx1:]...)
	skyline = append(skyline, skyline2[idx2:]...)

	return skyline
}

type PriorityQueue struct {
	value int
	next  *PriorityQueue
	head  *PriorityQueue
}

func (this *PriorityQueue) Insert(v int) {
	tmp := &PriorityQueue{value: v}
	if this.head.next == nil {
		this.head.next = tmp
		return
	}
	if this.head.next.value < v {
		tmp.next = this.head.next
		this.head.next = tmp
	} else {
		cur := this.head.next
		for cur.next != nil {
			if cur.next.value >= v {
				cur = cur.next
			} else {
				break
			}
		}
		tmp.next = cur.next
		cur.next = tmp
	}

}
func (this *PriorityQueue) Pop() {

	this.head = this.head.next
}
func (this *PriorityQueue) Remove(v int) {
	cur := this.head
	for cur.next != nil {
		if cur.next.value == v {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (this *PriorityQueue) Peek() int {
	if this.head.next == nil {
		return 0
	}
	return this.head.next.value
}
