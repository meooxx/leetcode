package main

func main() {}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// 1: loop1
// 1-1'-2-2'-3-3'
// 2:loop2 copy.Random = cur.random.Next
// |         |
// -----------
///  |         |
//   -----------
// loop3:1-2-3, 1'-2'-3'
func copyRandomList(head *Node) *Node {
	cur := head
	for cur != nil {
		n := &Node{
			Val: cur.Val,
		}
		next := cur.Next
		cur.Next = n
		n.Next = next
		cur = next
	}
	cur = head

	for cur != nil {
		curClone := cur.Next
		if cur.Random != nil {
			curClone.Random = cur.Random.Next
		}
		cur = curClone.Next
	}
	cur = head
	headClone := &Node{}
	last := headClone
	for cur != nil {
		curClone := cur.Next
		next := curClone.Next
		cur.Next = next
		last.Next = curClone
		last = last.Next
		cur = next
	}

	return headClone
}
