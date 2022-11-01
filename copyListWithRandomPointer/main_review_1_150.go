package main

type Node struct {
	Next   *Node
	Val    int
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	for cur != nil {
		next := cur.Next
		copyedNode := &Node{
			Val:  cur.Val,
			Next: next,
		}
		cur.Next = copyedNode
		cur = next
	}
	//1 -> 2  ==> 1-1'-2-2'
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head
	result := &Node{}
	pre := result
	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next

	}
	return result.Next

}
