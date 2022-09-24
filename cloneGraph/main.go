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
	return clone(node, &mp)
}

func clone(node *Node, m *map[int]*Node) *Node {
	mp := *m
	if node == nil {
		return nil
	}
	n := &Node{
		Val:       node.Val,
        Neighbors: make([]*Node,len(node.Neighbors)),
	}
	mp[node.Val] = n
	for i, nei := range node.Neighbors {

		if v, ok := mp[nei.Val]; ok {
			n.Neighbors[i] = v
		} else {
			n.Neighbors[i] = clone(nei, m)
		}
	}
	return n
}
