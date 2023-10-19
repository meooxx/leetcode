package main

import "sort"

func findItinerary(tickets [][]string) []string {
	graph := map[string][]string{}

	for i := range tickets {
		graph[tickets[i][0]] = append(graph[tickets[i][0]], tickets[i][1])
	}
	for i := range graph {
		sort.Slice(graph[i], func(a, b int) bool {
			return graph[i][a] < graph[i][b]
		})
	}
	anwser := []string{}
	dfs(&anwser, "JFK", &graph)
	for start, end := 0, len(anwser)-1; start < end; {
		if start < end {
			anwser[start], anwser[end] = anwser[end], anwser[start]
		}
		start++
		end--
	}
	return anwser
}

func dfs(anwser *[]string, start string, graph *map[string][]string) {
  // visit all edge
	for len((*graph)[start]) > 0 {
		dest := (*graph)[start]
		next := dest[0]
		dest = dest[1:]
		(*graph)[start] = dest
		dfs(anwser, next, graph)
	}
	// put the end point into anwser
	*anwser = append(*anwser, start)

}
