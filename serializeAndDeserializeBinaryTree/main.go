package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	vals []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.vals = append(this.vals, "#")
		return strings.Join(this.vals, ",")
	}
	this.vals = append(this.vals, strconv.Itoa(root.Val))
	this.serialize(root.Left)
	this.serialize(root.Right)
	return strings.Join(this.vals, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	var deserializeImpl func() *TreeNode
	deserializeImpl = func() *TreeNode {

		v := nodes[0]
		nodes = nodes[1:]
		if v == "#" {
			return nil
		}
		n, _ := strconv.Atoi(v)
		root := &TreeNode{
			Val: n,
		}
		root.Left = deserializeImpl()
		root.Right = deserializeImpl()
		return root
	}
	return deserializeImpl()

}

/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor();
* deser := Constructor();
* data := ser.serialize(root);
* ans := deser.deserialize(data);
 */
