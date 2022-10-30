package main

func main() {}

type Node struct {
	Val       int
	Neighbors []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	mp := map[int]*Node{}
	var clone func(n *Node) *Node
	clone = func(n *Node) *Node {
		if n == nil {
			return n
		}
		copyedNode := &Node{
			Val:       n.Val,
			Neighbors: make([]*Node, len(n.Neighbors)),
		}
		mp[n.Val] = copyedNode
		for i, n := range n.Neighbors {
			if v, ok := mp[n.Val]; ok {
				copyedNode.Neighbors[i] = v
			} else {
				copyedNode.Neighbors[i] = clone(n)
			}

		}
		return copyedNode
	}

	return clone(node)
}
