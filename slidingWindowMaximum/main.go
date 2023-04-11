package main

func maxSlidingWindow(nums []int, k int) []int {
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
