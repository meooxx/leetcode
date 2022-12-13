package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	mpEdge := map[int]int{}   // edges
	mpNode := map[int][]int{} // dependences
	stack := []int{}          // empty indegree node
	for i := range prerequisites {
		first := prerequisites[i][0]
		second := prerequisites[i][1]
		mpEdge[second]++
		if _, ok := mpEdge[first]; !ok {
			mpEdge[first] = 0
		}
		mpNode[first] = append(mpNode[first], second)
	}

	for key, v := range mpEdge {
		if v == 0 {
			stack = append(stack, key)
		}
	}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		for i := range mpNode[cur] {
			next := mpNode[cur][i]
			mpEdge[next]--
			if mpEdge[next] == 0 {
				stack = append(stack, next)
			}

		}
	}
	for _, v := range mpEdge {
		if v > 0 {
			return false
		}
	}
	return true

}
