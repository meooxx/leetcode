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

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	insertionSortList(head)
}
func insertionSortList(head *ListNode) *ListNode {
	result := &ListNode{Next: head}
	cur := head
	// 1 - 3 - 2 - 4
	// pre cur
	// 1 - 3 - 4	  cur.Next= cur.Next.Next
	// 2 - 3 - 4 	  cur.Next.Next = pre.Next
	// 1 - 2 - 3 - 4  pre.Next = cur.Next
	for cur != nil && cur.Next != nil {
		if cur.Val <= cur.Next.Val {
			cur = cur.Next
		} else {
			pre := result
			for pre.Next.Val < cur.Next.Val {
				pre = pre.Next
			}
			// 4 3 2 1
			next := cur.Next
			cur.Next = cur.Next.Next
			next.Next = pre.Next
			pre.Next = next
		}
	}
	return result.Next
}
