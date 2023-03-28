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
func isPalindrome(head *ListNode) bool {
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
