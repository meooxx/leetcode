package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: -4,
				},
			},
		},
	}
	head.Next.Next.Next.Next = head.Next
	fmt.Println(detectCycle(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// a: stand for distance form head - loophead(1-3)
// b: loophead to the meeting point of fast and slow(3-6)
// c: remaining of length of loop if fast and slow meets(6-3)
// 		a+2b+c == flow
// 		a+b == slow
// if fast meet slow and  fast twice speed slow
// 	then a+2b+c = 2(a+b) ==  c = a
//
//  a     b
// |   |   	 |
// 1 2 3 4 5 6 7
//     ^     ->
//     |------|
func detectCycle(head *ListNode) *ListNode {

	slow := head
	fast := head
	loopHead := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for loopHead != slow {
				slow = slow.Next
				loopHead = loopHead.Next
			}
			return loopHead
		}

	}
	return nil
}
