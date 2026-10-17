package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(list []int) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for _, n := range list {
		pre.Next = &ListNode{
			Val: n,
		}
		pre = pre.Next
	}
	return dummy.Next

}
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if right == left {
		return head
	}
	count := 0
	dummy := &ListNode{Next: head}
	pre := dummy
	// 1 23 45   2 4
	for count < left-1 {
		pre = pre.Next
		count++
	}
	tail := pre.Next
	for left < right {
		next := pre.Next
		pre.Next = tail.Next
		tail.Next = tail.Next.Next
		pre.Next.Next = next

		left++

	}
	return dummy.Next
}

func main() {
	reverseBetween(CreateListNode([]int{1, 2, 3, 4, 5}), 2, 3)
}
