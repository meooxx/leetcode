package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	var newl = &ListNode{}
// 	var rest, sum int = 0, 0
// 	var EMPT = &ListNode{}

// 	next1, next2 := l1, l2
// 	nextP := newl

// 	for next1 != nil || next2 != nil {

// 		if next1 == nil {
// 			next1 = EMPT
// 		}

// 		if next2 == nil {
// 			next2 = EMPT
// 		}

// 		n := next1.Val + next2.Val + rest
// 		sum = n % 10
// 		rest = n / 10

// 		nextP.Val = sum
// 		nextP.Next = &ListNode{
// 			Val:  rest,
// 			Next: nil}

// 		next1, next2 = next1.Next, next2.Next
// 		if next1 == nil && next2 == nil && rest <= 0 {
// 			nextP.Next = nil
// 		}

// 		nextP = nextP.Next

// 	}

// 	return newl

// }
