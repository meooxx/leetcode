package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
	reorderList(head)
	for cur := head; cur != nil; {
		fmt.Println(cur)

		cur = cur.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//   slow     fast
// 123      456  -> 123 654
// 1 用fast slow 找到中间节点
// 2 将slow 即中间的后面node翻转
// 3 slow后面节点依次将插入head
func reorderList(head *ListNode) {
	fast := head
	slow := head
	// 12345   3
	// 1234    2
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	cur := slow.Next
	var rightHead *ListNode = cur
	// revers the rightHead
	// 123 4 5 6 -> 123 654
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = rightHead
		rightHead = next
	}
	slow.Next = nil
	cur = head
	// 1-n-2-n-1
	for rightHead != nil {
		next := cur.Next
		cur.Next = rightHead
		rightHead = rightHead.Next
		cur.Next.Next = next

		cur = next

	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// passed
func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	cur := head
	nodes := []*ListNode{}
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}
	left := 1
	right := len(nodes) - 1
	// 1, 2, 3
	// 1, 2, 3, 4
	cur = head
	for left < right {
		next := cur.Next
		cur.Next = nodes[right]
		cur.Next.Next = next

		cur = next
		left++
		right--
	}
	nodes[right].Next = nil

}
