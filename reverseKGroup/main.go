package main

import "fmt"

func main() {

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	list := reverseKGroup(l1, 2)
	for list != nil {
		fmt.Println(list)
		list = list.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}

	count := 1

	//   last.next
	//      a           b
	//  	1   		2    3  4  5 ==>   2  1  3  4
	//
	//  	2   		1    3  4   ==>   3  2  1  4
	//  	3   		2    1  4
	//
	result := &ListNode{Next: head}
	lastNode := result
	tail := lastNode.Next
	next := tail.Next

	for {
		count = 0
		for tail != nil {
			tail = tail.Next
			count++
		}
		if count < k {
			break
		}

		count = 1
		tail = lastNode.Next
		for count < k && next != nil {

			nextNext := next.Next
			next.Next = lastNode.Next
			lastNode.Next = next
			tail.Next = nextNext

			next = nextNext
			count++
		}
		if count < k {
			break
		}

		lastNode = tail
		tail = lastNode.Next
		if tail == nil {
			break
		}
		next = tail.Next
	}
	return result.Next
}
