package main

import (
	"fmt"
)

/**
 * temp := &ListNode{}
 *
 * (if) l1 != nil  && l2 != nil ,
 * 		  compare (l1.Val, l2.Val) set temp.Val = l.Val(l which has smaller Val)
 * 			set ListNode = ListNode.Next which Val is smaller
 * 		this loop end
 * 		(if result = nil )at first loop set the result = temp
 * 		(if l1 != nil && l2 !-nil) set temp.Next = &ListNode{} (alloc new empty listNode)
 * 		(else) set temp.Next = l l!=nil
 * 	loops end
 * (else) set result = l l!=nil
 *
 *
 */
func main() {
	l := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 4}
	l.Next = l2
	l2.Next = l3

	n := &ListNode{Val: 1}
	n2 := &ListNode{Val: 3}
	// n3 := &ListNode{Val: 4}
	n.Next = n2
	// n2.Next = n3

	head := mergeTwoLists(l, n)
	for head != nil {
		fmt.Println(head, "head")
		head = head.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {

	nextL1 := l1
	nextL2 := l2

	result := &ListNode{}
	tmpList := result
	for nextL1 != nil && nextL2 != nil {
		if nextL1.Val < nextL2.Val {
			tmpList.Next = &ListNode{Val: nextL1.Val}
			nextL1 = nextL1.Next
			tmpList = tmpList.Next
		} else {
			tmpList.Next = &ListNode{Val: nextL2.Val}
			nextL2 = nextL2.Next
			tmpList = tmpList.Next
		}
	}
	if nextL1 != nil {
		tmpList.Next = nextL1

	}
	if nextL2 != nil {
		tmpList.Next = nextL2

	}

	return result.Next

}
