package main

import "fmt"

func main() {
	fmt.Println(canFinish0(5, [][]int{
		{0, 1}, {1, 2}, {1, 3}, {3, 4},
	}))
	fmt.Println(canFinish0(5, [][]int{
		{0, 1}, {1, 2}, {1, 3}, {3, 0},
	}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	mpEdge := map[int]int{}
	mpDep := map[int][]int{}
	stack := []int{}
	for i := range prerequisites {
		first := prerequisites[i][0]
		second := prerequisites[i][1]
		mpEdge[second]++
		mpDep[first] = append(mpDep[first], second)
	}
	for i := 0; i < numCourses; i++ {
		if mpEdge[i] == 0 {
			stack = append(stack, i)
		}
	}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		for _, v := range mpDep[cur] {
			mpEdge[v]--
			if mpEdge[v] == 0 {
				stack = append(stack, v)
			}
		}
	}
	for i := range mpEdge {
		if mpEdge[i] > 0 {
			return false
		}
	}
	return true
}

func canFinish0(numCourses int, prerequisites [][]int) bool {
	mpDep := map[int][]int{}

	for _, node := range prerequisites {
		mpDep[node[0]] = append(mpDep[node[0]], node[1])
	}
	visited := map[int]int{}
	var visit func(start int) bool
	visit = func(start int) bool {
		if visited[start] == 1 {
			return true
		}
		if visited[start] == 0 {
			visited[start] = 1
			deps := mpDep[start]
			for _, v := range deps {
				if visit(v) {
					return true
				}
			}
		}
		visited[start] = 2

		return false
	}
	for i := 0; i < numCourses; i++ {
		if visit(i) {
			return false
		}
	}
	return true
}
