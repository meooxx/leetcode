package main

import (
	"fmt"
	"strconv"
)

func main() {
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val: 2,
	}
	// l3 := &ListNode{
	// 	Val: 3,
	// }
	// l4 := &ListNode{
	// 	Val: 4,
	// }
	l1.Next = l2
	// l2.Next = l3
	// l3.Next = l4

	l := reverseBetween2(l1, 1, 2)
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	result := &ListNode{
		Next: head,
	}
	last := result
	cur := last.Next
	index := 1
	for cur != nil {
		if left == index {
			break
		}
		index++
		last = cur
		cur = cur.Next
	}

	// last     cur next
	// 1    	3   2    4 nil
	// right-1 实际交换了right 所以 < right
	for index < right {

		rest := last.Next
		next := cur.Next
		nextNext := next.Next

		last.Next = next
		next.Next = rest

		cur.Next = nextNext

		index += 1
	}
	return result.Next
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	result := &ListNode{
		Next: head,
	}
	pre := result
	cur := pre.Next
	count := 0
	tail := cur
	for cur != nil {
		tail = cur
		count++
		if count == left {
			break
		}
		pre = cur
		cur = cur.Next
	}

	var n *ListNode
	for cur != nil && count <= right {
		next := cur.Next
		cur.Next = n
		n = cur
		cur = next
		count++
	}

	pre.Next = n
	if cur != nil {
		tail.Next = cur
	}
	return result.Next
}


