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

// [1,2] [1,3], [3,4], [3,1]
//  1[2,3]
//  3[1,4] is cycled
// 1 visit[1] == 0 
// 		-> visit[1] = 1
// 		-> visit[2],   visit[3] 
//    	 visit[2]=2
//                   visit[3] => visit[1] == 1 => true , iscycle

func canFinish1(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	for i := range prerequisites {
		first := prerequisites[i][0]
		second := prerequisites[i][1]
		graph[first] = append(graph[first], second)
	}
	var isCycle func(graph map[int][]int, visit map[int]int, id int) bool
	isCycle = func(graph map[int][]int, visit map[int]int, id int) bool {
		if visit[id] == 1 {
			return true
		}
		if visit[id] == 0 {
			visit[id] = 1
			for i := range graph[id] {
				if isCycle(graph, visit, graph[id][i]) {
					return true
				}
			}
		}
		visit[id] = 2
		return false
	}
	visited := map[int]int{}
	for i := 0; i < numCourses; i++ {
		if isCycle(graph, visited, i) {
			return false
		}
	}
	return true

}
