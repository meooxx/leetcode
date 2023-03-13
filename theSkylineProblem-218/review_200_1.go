package main

import "sort"

func main() {
	// println(getSkyline([][]int{
	// 	{2, 9, 10},
	// 	{3, 7, 15},
	// 	{5, 12, 12},
	// 	{15, 20, 10},
	// 	{19, 24, 8},
	// }))
	println(getSkyline([][]int{
		{2, 5, 10},
		{5, 7, 10},

	}))
}

type Queue struct {
	Val  int
	Next *Queue
}

func (this *Queue) Insert(v int) {
	if this.Next == nil {
		this.Next = &Queue{Val: v}
	} else {
		cur := this
		newNode := &Queue{Val: v}
		// 5 3
		for cur.Next != nil {
			if cur.Next.Val < v {
				next := cur.Next
				cur.Next = newNode
				cur.Next.Next = next
				return
			} else {
				cur = cur.Next
			}
		}
		if cur.Next == nil {
			cur.Next = newNode
		}
	}

}

func (this *Queue) Remove(v int) {
	cur := this
	for cur.Next != nil {
		if cur.Next.Val == v {
			cur.Next = cur.Next.Next
			return
		} else {
			cur = cur.Next
		}
	}
}

func (this *Queue) Peek() int {
	if this.Next != nil {
		return this.Next.Val
	}
	return 0
}

func getSkyline(buildings [][]int) [][]int {
	nodes := [][]int{}
	for _, v := range buildings {
		nodes = append(nodes, []int{
			v[0], -v[2],
		}, []int{v[1], v[2]})
	}
	sort.Slice(nodes, func(a, b int) bool {
		if nodes[a][0] != nodes[b][0] {
			return nodes[a][0] <= nodes[b][0]
		}
		return nodes[a][1] <= nodes[b][1]

	})
	result := [][]int{}
	q := Queue{}
	pre := 0
	for _, v := range nodes {
		if v[1] < 0 {
			q.Insert(-v[1])
		} else {
			q.Remove(v[1])
		}
		maxHeight := q.Peek()
		if maxHeight != pre {
			pre = maxHeight
			result = append(result, []int{v[0], maxHeight})
		}
	}
	return result

}
