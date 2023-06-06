package main

func findMinHeightTrees1(n int, edges [][]int) []int {
	mp := make([][]int, n)

	for _, v := range edges {
		mp[v[0]] = append(mp[v[0]], v[1])
		mp[v[1]] = append(mp[v[1]], v[0])
	}
	result := []int{}
	maxLevel := n
	for i := 0; i < n; i++ {
		level := -1
		q := []int{i}
		visited := map[int]bool{}
		for len(q) > 0 {
			level++
			if level > maxLevel {
				continue
			}
			qSize := len(q)
			for j := 0; j < qSize; j++ {
				cur := q[0]
				q = q[1:]
				if visited[cur] {
					continue
				}
				visited[cur] = true
				q = append(q, mp[cur]...)
			}
		}
		if level <= maxLevel {
			if level < maxLevel {
				result = []int{}
				maxLevel = level
			}
			result = append(result, i)
		}
	}
	return result
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	mp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		mp[i] = map[int]int{}
	}

	for _, v := range edges {
		mp[v[0]][v[1]] = 1
		mp[v[1]][v[0]] = 1
	}

	leaves := []int{}
	for i := range mp {
		if len(mp[i]) == 1 {
			leaves = append(leaves, i)
		}
	}
	for n > 2 {
		// 1-3-2  (remove 1,2)  3
		size := len(leaves)
		n -= size
		for i := 0; i < size; i++ {
			leave := leaves[0]
			leaves = leaves[1:]
			// leaves that only have one parent
			for key := range mp[leave] {
				delete(mp[key], leave)
				if len(mp[key]) == 1 {
					leaves = append(leaves, key)
				}
			}
		}
	}
	return leaves

}
