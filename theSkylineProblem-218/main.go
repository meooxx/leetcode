package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
	fmt.Println(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}))

}

func getSkyline(buildings [][]int) [][]int {
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
