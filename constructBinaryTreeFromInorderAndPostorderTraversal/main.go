package main
import "fmt"
func main(){
	n := buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3})
	fmt.Println(n)
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
    Val int
    Left *TreeNode
    Right *TreeNode
}
 func buildTree(inorder []int, postorder []int) *TreeNode {
    mp:= map[int]int{}
    for  index,v:= range  inorder  {
        mp[v] = index
    }
    n := getRoot(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1, &mp)
    return n
}

func getRoot(inorder, postorder []int, inStart, inEnd, postStart, postEnd int, m *map[int]int) *TreeNode{
    if postStart > postEnd {
        return nil
    }
    mp := *m
    root := postorder[postEnd]
    rootIndex := mp[root]
    
    r := &TreeNode{Val:root}
    leftNum  := rootIndex - inStart
	// [9,3,15,20,7] inStart, inEnd
	// [9,15,7,20,3] postStart
    left:= getRoot(inorder, postorder, inStart, rootIndex-1, postStart, postStart+leftNum-1, m)
    
    right := getRoot(inorder, postorder, rootIndex+1, inEnd, postStart+leftNum, postEnd-1, m)
    
    r.Left = left
    r.Right = right
    return r
}
