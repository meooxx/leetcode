package main

// graph   1:[2,3]
// degree  2:1, 3: 1
// find O degree node == [1],cuz degree[1] == 0, push(1)
// loop  graph[1]: 
//    graph[i] == [2,3]
// 					degree[2]--, degree[2] == 0, push(2)
// 					degree[3]--, degree[3] == 0, push(3)

// result 1 2 3 => reverse 3 2 1

func findOrder(numCourses int, prerequisites [][]int) []int {
	// [3,1] [2,1]
	// 3: [1]
	// 2: [1]
	graph := map[int][]int{}
	indegree := map[int]int{}
	for _, v := range prerequisites {
		first := v[0]
		second := v[1]
		graph[first] = append(graph[first], second)
		indegree[second]++
	}
	result := []int{}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		cur := q[0]
		result = append(result, cur)
		q = q[1:]
		for _, v := range graph[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(result) < numCourses-1 {
		return []int{}
	}
	for left, right := 0, len(result)-1; left < right; {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result

}

// graph   2:[1], 3:[1]
// degree  1:2
// find O degree node == [2,3],cuz degree[2] ==0 and degree[3] == 0
// loop i graph[2,3]: 
// 		graph[2] == [1] => degree[1]--
//  	graph[3] == [1] => degree[1]--
// degree[1] == 0, push(1)

func findOrder1(numCourses int, prerequisites [][]int) []int {
	graph := map[int][]int{}
	degree := map[int]int{}
	for _, v := range prerequisites {
		first := v[0]
		second := v[1]
		graph[second] = append(graph[second], first)
		degree[first]++
	}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			q = append(q, i)
		}
	}
	result := []int{}
	for len(q) > 0 {
		cur := q[0]
		result = append(result, cur)
		q = q[1:]
		for _, v := range graph[cur] {
			degree[v]--
			if degree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(result) < numCourses - 1 {
		return []int{}
	}
	return result

}
