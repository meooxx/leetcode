package leetcodetypes

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
