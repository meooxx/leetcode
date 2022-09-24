package main

func main() {
	root := &Node{
		Val: 3,
		Left: &Node{
			Val: 9,
		},
		Right: &Node{
			Val: 20,
			Left: &Node{
				Val: 15,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	connect(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
[1,2,3,4,5,null,6,7,null,null,null,null,8]
           1
       2       3
     4  5     nil 6
	7 nil nil nil nil 8
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	head := root
	// tmpNode
	// last == tmpNode,当遇到第一个节点时候, tmpNode.Next设置成第一个有效节点
	// 随后 last 设置成 第一个有效节点
	// tmpNode Next 中保留第一个节点的引用
	tmpNode := &Node{}
	for head != nil {
		cur := head
		last := tmpNode
		for cur != nil {
			if cur.Left != nil {
				last.Next = cur.Left
				last = cur.Left
			}
			if cur.Right != nil {
				last.Next = cur.Right
				last = cur.Right
			}
			cur = cur.Next
		}
		head = tmpNode.Next
		tmpNode.Next = nil
	}
	return root

}
