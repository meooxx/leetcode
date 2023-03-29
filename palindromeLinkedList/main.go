package main

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
func isPalindrome1(head *ListNode) bool {
	nodes := []int{}
	cur := head
	for cur != nil {
		nodes = append(nodes, cur.Val)
		cur = cur.Next
	}
	for i, j := 0, len(nodes)-1; i < j; {
		if nodes[i] != nodes[j] {
			return false
		}
		i++
		j--
	}
	return true
}
// space: O(N)
func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 1234, 12345
	//  s      s
	backHalfHead := slow.Next
	slow.Next = nil
	cur := backHalfHead
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = backHalfHead

		backHalfHead = next
	}

	for head != nil && backHalfHead != nil {
		if head.Val != backHalfHead.Val {
			return false
		}
		head = head.Next
		backHalfHead = backHalfHead.Next
	}
	return true

}
