package main

import "math/rand"

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
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

// 1 -> 2 -> 3 -> 4
// anwser  =  0
//            1, 1/[0, 1) anwser = 1, P(1) = 1
//            2, 1/[0, 2) anwser = 2, P(2) = 1/2, P(1) = 1-1/2
//            3, 1/[0, 3) anwser = 3, P(3) = 1/3, P(2) = 1/2 * (1-1/3), P(1) = 1/3
func (this *Solution) GetRandom() int {
	anwser := 0
	n := 0
	cur := this.head
	for cur != nil {
		n++
		r := rand.Intn(n)
		if r == n-1 {
			anwser = cur.Val
		}
		cur = cur.Next
	}
	return anwser
}

/**
* Your Solution object will be instantiated and called as such:
* obj := Constructor(head);
* param_1 := obj.GetRandom();
 */
