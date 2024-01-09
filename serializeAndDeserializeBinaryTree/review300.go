package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
	vals strings.Builder
}

func Constructor() Codec {
	return Codec{
		vals: strings.Builder{},
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.vals.WriteByte('#')
		this.vals.WriteByte(',')
		return this.vals.String()
	}
	this.vals.WriteString(strconv.Itoa(root.Val))
	this.vals.WriteByte(',')

	// recursively serialize root.Left first, and deserialize root.Left first
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.vals.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tn := &TreeNode{}
	nodes := strings.Split(data, ",")
	if len(nodes) == 0 {
		return tn
	}
	var deserializeImpl func() *TreeNode
	deserializeImpl = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}
		b := nodes[0]
		nodes = nodes[1:]

		if b == "#" {
			return nil
		}
		v, _ := strconv.Atoi(string(b))
		t := &TreeNode{
			Val: v,
		}
		t.Left = deserializeImpl()
		t.Right = deserializeImpl()
		return t
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
