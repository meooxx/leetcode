package main
import "fmt"
func main() {
	r := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isBalanced(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
    Abs:=func (a int) int{
        if a < 0 {
            return -a
        }
        return a
    }
    return Abs(getMaxDepth(root.Left) - getMaxDepth(root.Right)) <= 1 && isBalanced(root.Left) &&isBalanced(root.Right)
}

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1
	right := 1
	if root.Left != nil {
		left += getMaxDepth(root.Left)
	}
	if root.Right != nil {
		right += getMaxDepth(root.Right)
	}
	if left >= right {
		return left
	}
	return right
}
