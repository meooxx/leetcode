package main


func main(){

}


type TreeNode struct {
    Val int
    Left *TreeNode
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
func minDepth(root *TreeNode) int {
    return getMinDepth(root)
}

func getMinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1
	right := 1
	left+= getMinDepth(root.Left)
	right += getMinDepth(root.Right)

    if root.Left == nil {
        return right
    }
    if root.Right == nil {
        return left
    }
    
	if left <= right {
		return left
	}
	return right

}
