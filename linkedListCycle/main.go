package main

func main() {}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    if head == nil {
        return false 
    }
    fast:= head.Next
    slow := head
    
    for slow != nil && fast != nil {
        if fast == slow {
            return true
        }
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
        slow = slow.Next
    }
    return false
}
