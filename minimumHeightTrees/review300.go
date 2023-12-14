package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([]map[int]int, n)

	for i := range graph {
		graph[i] = map[int]int{}
	}
	// [a, b], a or b can be root
	for _, v := range edges {
		graph[v[0]][v[1]] = 1
		graph[v[1]][v[0]] = 1
	}

	leaves := []int{}

	for key := range graph {
		if len(graph[key]) == 1 {
			leaves = append(leaves, key)
		}
	}

	for n > 2 {
		qSize := len(leaves)
		n -= qSize
		for i := 0; i < qSize; i++ {
			node := leaves[0]
			leaves = leaves[1:]
			// root
			//  a
			// b c
			for n := range graph[node] {
				// e.g. delete c from a
				delete(graph[n], node)
				// a-b, a-c, a-root
				// if len(graph[a]) == 1, that means a a leave node
				// only a-root, if graph[a] == 0, that means a has no edges
				if len(graph[n]) == 1 {
					leaves = append(leaves, n)
				}
			}
		}
	}
	return leaves
}
