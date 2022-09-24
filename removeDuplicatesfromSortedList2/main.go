package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val: 1,
	}
	l3 := &ListNode{
		Val: 1,
	}
	l4 := &ListNode{
		Val: 1,
	}
	l5 := &ListNode{
		Val: 1,
	}
	l6 := &ListNode{
		Val: 1,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	h := deleteDuplicates(l1)


	for h != nil {
		fmt.Println(h)
		h = h.Next
	}

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
	result := &ListNode{
		Next: head,
	}
	//cur=> 1=>1=>1=>2=>2=>3
	last := result
	cur := last.Next

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			next := cur.Next
			for next != nil {
				if cur.Val == next.Val {
					next = next.Next
				} else {
					break
				}
			}
			cur = next
			last.Next = cur
		} else {
			last = cur
			cur = last.Next

		}

	}
	return result.Next
}
