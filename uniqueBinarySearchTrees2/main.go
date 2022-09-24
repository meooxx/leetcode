package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Printf("%#v", generateTrees(3))
	v, _ := json.MarshalIndent(generateTrees2(3), "", "  ")

	fmt.Printf("%s", string(v))

}

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

func generateTrees(n int) []*TreeNode {

	return getTreeNodes(1, n)
}

func getTreeNodes(start, end int) []*TreeNode {
	result := []*TreeNode{}

	if start > end {
		return []*TreeNode{nil}
	}
	for i := start; i <= end; i++ {
		leftNodes := getTreeNodes(start, i-1)
		rightNodes := getTreeNodes(i+1, end)

		for _, left := range leftNodes {
			for _, right := range rightNodes {
				root := &TreeNode{Val: i, Left: left, Right: right}
				result = append(result, root)
			}
		}

	}

	return result

}
// dp[0] = [nil] base
// dp[1] = [1]
// dp[2] = root==1:[nil] * d[1]+offset || root==2:dp[1] * [nil]

//  n 拆成例如 1, 2, 3,4
// 把1设为root [nil] * dp[3]
// 设2位root:left * right ==   dp[1] * d[2]
// 左侧节点已经是算好的 右侧dp[2]表示 (3,4)俩个数字这个组合
// 即可使用dp[2]的每个结果加上偏移值可得到 (3,4)组合的结果, 偏移值即root 1-2 -> 3-4, 2-1 -> 4-3
// 设3为root: left * right == dp[3]*dp[1]+offset 右侧dp[1]要加上3才能获得准确的结果
// 设4为root: dp[3]* [nil]
// 所有可能汇总
func generateTrees2(n int) []*TreeNode {

	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{nil}

	for target := 1; target <= n; target++ {
		dp[target] = []*TreeNode{}
		for root := 1; root <= target; root++ {
			left := root - 1
			right := target - root

			for _, leftNode := range dp[left] {
				for _, rightNode := range dp[right] {
					node := &TreeNode{
						Val:   root,
						Left:  leftNode,
						Right: copyNode(rightNode, root),
					}
					dp[target] = append(dp[target], node)
				}
			}

		}
	}
	return dp[n]
}

func copyNode(n *TreeNode, offset int) *TreeNode {
	if n == nil {
		return nil
	}
	return &TreeNode{
		Val:   n.Val + offset,
		Left:  copyNode(n.Left, offset),
		Right: copyNode(n.Right, offset),
	}
}
