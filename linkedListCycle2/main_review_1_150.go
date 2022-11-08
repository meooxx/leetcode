package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//   h + 2meet   =  h + meet + rest
// meet = rest
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			cur := head
			for cur != slow {
				cur = cur.Next
				slow = slow.Next
			}
			return slow
		}

	}
	return nil
}
