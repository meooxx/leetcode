package main

func main() {}

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

/*
 * Given 1 -> 3 -> 2 -> 4 - > null
 
    result -> 1 -> 3 -> 2 -> 4 - > null
                   |    |
             3>2  cur toInsert
				  	
    result -> 1 -> 3 -> 2 -> 4 - > null
	 ^	
	 |
	step1 .next 1<2✅		   
	    step2 .Next 3<2❌2,  preInsert==node1
 *
 *
 *
*/
// it cost less runtime
func insertionSortList(head *ListNode) *ListNode {
	result := &ListNode{
		Next: head,
	}
	for cur := result.Next; cur.Next != nil; {
		if cur.Val <= cur.Next.Val {
			cur = cur.Next
		} else {
			node := cur.Next
			preInsert := result
			for preInsert.Next.Val <= node.Val {
				preInsert = preInsert.Next
			}
			cur.Next = node.Next
			node.Next = preInsert.Next
			preInsert.Next = node

		}
	}
	return result.Next
}

// passed
func insertionSortList2(head *ListNode) *ListNode {
	result := &ListNode{}
	for head != nil {
		cur := result
		for cur != nil && cur.Next != nil {
			if head.Val >= cur.Next.Val {
				cur = cur.Next
			} else {
				break
			}
		}
		next := cur.Next
		cur.Next = head
		head = head.Next
		cur.Next.Next = next
	}
	return result.Next

}
