package main

import "fmt"

// import "fmt"
func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))

}


func findRepeatedDnaSequences(s string) []string {
	result := []string{}
	mp := map[string]int{}

	for i := 0; i+9 < len(s); i++ {
		seg := s[i : i+10]
		mp[seg]++
		if mp[seg] == 2 {
			result = append(result, seg)
		}
	}

	return result
}
