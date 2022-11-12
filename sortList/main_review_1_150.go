package main

func main() {}

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

func sortList(head *ListNode) *ListNode {
	size := 0

	split := func(list *ListNode, n int) *ListNode {
		if list == nil {
			return list
		}
		for i := 1; list.Next != nil && i < n; i++ {
			list = list.Next
		}
		next := list.Next
		list.Next = nil
		return next
	}

	mergeList := func(left, right, pre *ListNode) *ListNode {
		for left != nil && right != nil {
			if left.Val <= right.Val {
				pre.Next = left
				pre = pre.Next
				left = left.Next
			} else {
				pre.Next = right
				pre = pre.Next
				right = right.Next
			}
		}
		if left != nil {
			pre.Next = left
		}
		if right != nil {
			pre.Next = right
		}
		for pre.Next != nil {
			pre = pre.Next
		}
		return pre
	}
	for cur := head; cur != nil; cur = cur.Next {
		size++
	}
	result := &ListNode{Next: head}
	for n := 1; n < size; n = n << 1 {
		cur := result.Next
		pre := result
		for cur != nil {
			left := cur
			right := split(left, n)
			rightEnd := split(right, n)

			pre = mergeList(left, right, pre)
			cur = rightEnd
		}
	}
	return result.Next

}

func merge(left, right *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left.Next
			left = left.Next
			cur = cur.Next
		} else {
			cur.Next = right
			right = right.Next
			cur = cur.Next
		}
	}
	if left != nil {
		cur.Next = right
	}
	if right != nil {
		cur.Next = right
	}
	return result.Next
}
func sortListRec(head *ListNode) *ListNode {
	fast := head
	slow := head
	// 1234
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil

	left := sortListRec(head)
	right := sortListRec(next)
	return merge(left, right)

}
