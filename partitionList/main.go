package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
	}
	l2 := &ListNode{
		Val: 2,
	}
	l3 := &ListNode{
		Val: 3,
	}
	l4 := &ListNode{
		Val: 4,
	}
	l5 := &ListNode{
		Val: 5,
	}
	l1.Next = l4
	l4.Next = l3
	l3.Next = l2
	l2.Next = l5

	l := partition(l1, 3)
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1, 4, 2,5
// lessXList   pre  curr   next
// 1  		   4      2    5
// 1->2->4->5

func partition1(head *ListNode, x int) *ListNode {
	result := &ListNode{
		Next: head,
	}

	lessX := result
	for lessX.Next != nil && lessX.Next.Val < x {
		lessX = lessX.Next
	}

	pre := lessX
	curr := pre.Next
	// 1,4,3,2,5
	for curr != nil {
		if curr.Val < x {
			next := curr.Next
			pre.Next = next
			rest := lessX.Next

			lessX.Next = curr
			curr.Next = rest
			lessX = lessX.Next

			// 3  [1]  2, curr==1, => [3],2
			// 方便最后俩行赋值
			// pre = curr == 3
			// cur = curr.Next == 3.Next == 2
			curr = pre
		}
		pre = curr
		curr = curr.Next
	}
	return result.Next
}

func partition(head *ListNode, x int) *ListNode {
	result := &ListNode{
		Next: head,
	}

    pre:= result
    tail:= pre
    cur := tail.Next
    for cur != nil {
        if cur.Val < x {
            pre.Next = cur.Next
            cur.Next = tail.Next
            tail.Next = cur
            
            tail = cur
            cur = pre.Next
            
        }else {
            pre = cur
            cur = cur.Next
        
        }
        
    }
    
    return result.Next
    
}
