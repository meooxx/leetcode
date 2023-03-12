package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	mpEdges := map[int]int{}
	mpNodes := map[int][]int{}

	for i := range prerequisites {
		node := prerequisites[i]
		mpEdges[node[0]]++
		mpNodes[node[1]] = append(mpNodes[node[1]], node[0])
	}

	stack := []int{}
	result := []int{}
	for i := 0; i < numCourses; i++ {
		if mpEdges[i] == 0 {
			stack = append(stack, i)
		}
	}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		result = append(result, cur)
		for _, next := range mpNodes[cur] {
			mpEdges[next]--
			if mpEdges[next] == 0 {
				stack = append(stack, next)
			}
		}
	}
	for i := range mpEdges {
		if mpEdges[i] > 0 {
			return []int{}
		}
	}
	return result
}
