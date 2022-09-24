package main
import "fmt"



/**
 * 1->2->3->4
 * 
 * 2 -1 -34
 * 
 * 3 - 21 -4
 * 
 * 4- 321-nil
 * 
 * temp = &ListNode{Val:curr.Val}
 * temp.Next = &ListNode{Val:preNode.Val, Next:preNode.Next}
 * rest = curr.Next
 * preNode = temp
 * tail = temp.Next
 * tail.Next = rest
 * 
 * count ++
 * curr = rest
 * 
 * 
 */



func main() {
	
	l1 := &ListNode{Val:1}
	l2 := &ListNode{Val:2}
	l3 := &ListNode{Val:3}
	l4 := &ListNode{Val:4}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	list := reverseKList(l1, 5)
	for list != nil {
		fmt.Println(list)
		list = list.Next
	}

}



type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKList(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	} 

	count := 1

	var preNode *ListNode = &ListNode{Val:head.Val}
	var tail *ListNode  = &ListNode{}
	var curr *ListNode = head.Next

	var rest *ListNode
	var temp *ListNode  = &ListNode{}

	for  count < k {
		rest = curr.Next
		temp = &ListNode{Val: curr.Val}
		temp.Next = &ListNode{Val:preNode.Val,Next: preNode.Next}

		if tail.Next == nil {
			tail = temp.Next
		}

		tail.Next = rest

		preNode = temp
		curr = rest
		count ++
		if curr == nil {
			break
		}
	}
	return  preNode

}