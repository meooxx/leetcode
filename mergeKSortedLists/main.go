package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 5} //
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 4} //

	t1 := &ListNode{Val: 2}
	t2 := &ListNode{Val: 6}
	t1.Next = t2

	l1.Next = l2
	l2.Next = l3
	//l2.Next = l3
	n1.Next = n2
	n2.Next = n3

	// c:= &ListNode{Val:2}
	// c1:= &ListNode{}
	// c2:= &ListNode{Val: -1}

	list := mergeKLists([]*ListNode{l1})
	for list != nil {
		fmt.Println(list)
		list = list.Next
	}

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	result := mergeTwoLists(mergeTwoLists(lists[0], lists[1]), mergeKLists(lists[2:]))
	return result

}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	nextL1, nextL2 := l1, l2
	for nextL1 != nil && nextL2 != nil {
		if nextL1.Val <= nextL2.Val {
			tmp.Next = &ListNode{Val: nextL1.Val}
			nextL1 = nextL1.Next
			tmp = tmp.Next
		} else {
			tmp.Next = &ListNode{Val: nextL2.Val}
			nextL2 = nextL2.Next
			tmp = tmp.Next
		}
	}
	if nextL1 != nil {
		tmp.Next = nextL1
	}
	if nextL2 != nil {
		tmp.Next = nextL2
	}
	return result.Next
}
