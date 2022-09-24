package main

func main() {}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
	pre := root
	for pre.Left != nil {
		cur := root

		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Left
	}
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	stack := []*Node{root}

	for len(stack) > 0 {
		s := []*Node{}
		var pre *Node

		for len(stack) > 0 {
			top := stack[0]
			if pre != nil {
				pre.Next = top
			}
			pre = top
			stack = stack[1:]

			if top.Left != nil {
				s = append(s, top.Left)
			}
			if top.Right != nil {
				s = append(s, top.Right)
			}

		}
		stack = s

	}
	return root
}
