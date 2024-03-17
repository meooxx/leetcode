package main

func findOrder(numCourses int, prerequisites [][]int) []int {

	graph := map[int][]int{}
	degree := map[int]int{}
	for _, node := range prerequisites {
		graph[node[1]] = append(graph[node[1]], node[0])
		degree[node[0]]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	result := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		for _, n := range graph[node] {
			degree[n]--
			if degree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}
	for i := range degree {
		if degree[i] > 0 {
			return []int{}
		}
	}
	return result
}
