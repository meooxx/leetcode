package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {

	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		size := len(stack)
		result = append(result, stack[size-1].Val)
		for i := 0; i < size; i++ {
			cur := stack[0]
			stack = stack[1:]
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			// 可以在开始前加入
			// if i == size - 1 {
			//    result = append(result, cur.Val)

			// }
		}
	}
	return result

}
