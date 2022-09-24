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
 func levelOrderBottom(root *TreeNode) [][]int {
    
    result := [][]int{}
    if root == nil {
        return result   
    }
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        s := []*TreeNode{}
        item := []int{}
        for _, v := range stack {
            item =  append(item, v.Val)
            if v.Left != nil {
                s = append(s, v.Left)
            }
            if v.Right != nil {
                s = append(s, v.Right)
            }
        }
        stack = s
       
        result = append(result, item)
    }
    for left, right:=0,len(result)-1; left < right; {
        
        result[left], result[right] = result[right], result[left]
        left++
        right--
    }
    return result
}