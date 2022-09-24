package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
	input1 := []string{"", "", "", "ate", "nat", "bat"}
	fmt.Printf("%#v", groupAnagrams(input1))
}

func groupAnagrams(strs []string) [][]string {
	answer := [][]string{}
	m := map[string][]string{}
	for _, v := range strs {
		strSlice := strings.Split(v, "")
		sort.Strings(strSlice)
		str := strings.Join(strSlice, "")
		m[str] = append(m[str], v)
	}
	for _, v := range m {
		answer = append(answer, v)
	}
	return answer
}
