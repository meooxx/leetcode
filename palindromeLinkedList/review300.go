package main

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

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head

	// 1->2
	// slow == 1, slow.Next == 2  
	// fast == 1
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 1 2 3 2 1
	//     ^
	//   slow
	midHead := slow.Next
	slow.Next = nil

	for cur:=midHead; cur!= nil && cur.Next != nil; {
		next := cur.Next
		cur.Next = next.Next
		next.Next = midHead
		midHead = next
	}

	for cur := midHead; cur != nil; cur = cur.Next {
		if head.Val != cur.Val {
			return false
		}
		head = head.Next
	}
	return true
}
