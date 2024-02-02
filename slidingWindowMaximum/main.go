package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{
		1, 3, -1, -3, 5, 3, 6, 7,
	}, 3))
}

func maxSlidingWindow1(nums []int, k int) []int {
	maxHeap := &MaxHeap{&Node{}}
	left, right := 0, 0
	result := []int{}
	for right < len(nums) {

		for right-left < k {
			maxHeap.push(nums[right])
			right++
		}
		result = append(result, maxHeap.head.Next.Val)
		maxHeap.pop(nums[left])
		left++
	}
	return result
}
/*
*   2 1 3 5 -1 2
*  push(2)  [2]
		 push(1) [2,1]
		3>1 pop(), 3>2 pop() -> [] ->  push(3)
		5>3	pop() push(5) -> [5]
		-1<5 push(-1) -> [5, -1]
		2 >-1<5, pop(), push(2) -> [5,2]
*/
func maxSlidingWindow(nums []int, k int) []int {
	q := []int{} // indices
	result := []int{}
	for i := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		// is same  as {i - k == q[0]}
		// 1 4 2 5
		//     if i-k == q[0], the i will be at  position '5'
    if i - k + 1 > q[0] {
      q = q[1:]
    }
		if i-k+1 >= 0 {
			result = append(result, nums[q[0]])
		}
	}
	return result
}

type Node struct {
	Val  int
	Next *Node
}

type MaxHeap struct {
	head *Node
	// map [int]*Node
}

func (this *MaxHeap) push(n int) {
	head := this.head

	cur := head
	for cur.Next != nil {
		if n > cur.Next.Val {
			next := cur.Next
			cur.Next = &Node{Val: n, Next: next}
			break
		} else {
			cur = cur.Next
		}
	}
	if cur.Next == nil {
		cur.Next = &Node{Val: n}
	}
}

func (this *MaxHeap) pop(n int) {
	cur := this.head

	for cur.Next != nil {
		if cur.Next.Val == n {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}

}
