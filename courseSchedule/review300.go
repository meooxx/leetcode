package main

func canFinish(numCourses int, prerequisites [][]int) bool {

	node := map[int][]int{}
	edges := map[int]int{}

	for _, v := range prerequisites {
		node[v[1]] = append(node[v[1]], v[0])
		edges[v[0]]++
	}
	stack := []int{}
	for i := 0; i < numCourses; i++ {
		if edges[i] == 0 {
			stack = append(stack, i)
		}
	}
	// 1 -> 0
	// 1 -> 2
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		for _, n := range node[cur] {
			edges[n]--
			if edges[n] == 0 {
				stack = append(stack, n)
			}
		}
	}

	for i := range edges {
		if edges[i] > 0 {
			return false
		}
	}
	return true

}
