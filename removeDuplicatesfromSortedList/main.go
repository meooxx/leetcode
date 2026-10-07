package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 1}
	l4 := &ListNode{Val: 2}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	fmt.Printf("%#v", deleteDuplicates(l1))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	answer := &ListNode{}
	pre := answer
	pre.Next = head
	curr := head
	for curr != nil {

    // find the last the duplicated node
		// 1 2 2  2    3
	  //        ^ CURR
		for curr.Next  != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
		}
		
		if pre.Next == curr {
			// curr remain unchanged
		// no duplicates for the curr node
			pre = pre.Next
		} else {
			// assume that next node is potential eligible node 
			pre.Next  = curr.Next
		}
		curr = curr.Next
	}
	return answer.Next
}
