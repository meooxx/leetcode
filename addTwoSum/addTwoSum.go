package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// l1 -> l2 -> l3
// ll1 -> ll2 ->ll3
// 1 => 2 =>3
// 9 => 1 =>1
// r: 0 => 4 => 4

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}

	n1 := l1
	n2 := l2
	nR := r
	sum, rest := 0, 0
	for n1 != nil || n2 != nil {
		if n1 == nil {
			n1 = &ListNode{}
		}
		if n2 == nil {
			n2 = &ListNode{}
		}
		sum = (n1.Val + n2.Val + nR.Val) % 10
		rest = (n1.Val + n2.Val + nR.Val) / 10
		nR.Val = sum
		if n1.Next != nil || n2.Next != nil || rest > 0 {
			nR.Next = &ListNode{
				Val: rest, // 暂存进位 5+6, 保存1
			}
		}

		n1, n2 = n1.Next, n2.Next
		nR = nR.Next
	}

	return r
}

func main() {
	l1 := &ListNode{Val: 4, Next: &ListNode{Val: 8, Next: &ListNode{Val: 6}}}
	l2 := &ListNode{Val: 7, Next: &ListNode{Val: 1}}

	r := addTwoNumbers(l1, l2)
	for r != nil {
		fmt.Printf("%v", r)
		r = r.Next
	}
}
