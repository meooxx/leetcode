package main

import (
	"fmt"
	"sort"
)
//      SFO
//       ^   ^
//       |    \ \
//              v
//    JFK ->   ATL
//        <-
//   
//  JFK -> ATL -> JFK -> SFO -> ATL -> SFO, add(SFO)
//          ....      add(SFO) add(ATL) 
// if out of ticket , add it to anwser it will be an end   
// 
//  JFK  -> SFO (❌, in order)
func main() {
	fmt.Println(findItinerary([][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}))
}

func findItinerary(tickets [][]string) []string {

	graph := map[string][]string{}
	for i := range tickets {
		item := tickets[i]
		start := item[0]
		dest := item[1]
		graph[start] = append(graph[start], dest)
	}
	for key := range graph {
		sort.Slice(graph[key], func(a, b int) bool {
			return graph[key][a] < graph[key][b]
		})
	}
	anwser := []string{}
	dfs(graph, "JFK", &anwser)
	for left, right := 0, len(anwser)-1; left < right; {
		anwser[left], anwser[right] = anwser[right], anwser[left]
		left++
		right--
	}
	return anwser

}

func dfs(graph map[string][]string, key string, anwser *[]string) {
	// try all ticket
	for len(graph[key]) > 0 {
		list := graph[key]
		nextKey := list[0]
		graph[key] = list[1:]
		list = list[1:]
		dfs(graph, nextKey, anwser)
	}
	// append the end destination 
	*anwser = append(*anwser, key)
}
