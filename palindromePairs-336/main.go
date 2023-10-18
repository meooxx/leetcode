package main

import "fmt"

func main() {
	// fmt.Println(palindromePairs([]string{
	// 	"ba","a" , "aaa",
	// }))
	fmt.Println(palindromePairs([]string{
		"bat","tab","cat",
	}))
}

type Trie struct {
	Index int
	Next  [26]*Trie
	List  []int
}

func palindromePairs(words []string) [][]int {
	root := &Trie{Index: -1, Next: [26]*Trie{}, List: []int{}}
	for i := range words {
		addWords(words[i], i, root)
	}

	result := [][]int{}
	for i := range words {
		search(words[i], i, root, &result)
	}
	return result
}

func addWords(word string, index int, root *Trie) {
	for i := len(word) - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		next := root.Next
		if next[idx] == nil {
			root.Next[idx] = &Trie{Index: -1, Next: [26]*Trie{}, List: []int{}}
		}
		if isPalindrome(word, 0, i) {
			root.List = append(root.List, index)
		}
		root = root.Next[idx]
	}
	root.Index = index
	root.List = append(root.List, index)
}

func search(word string, index int, root *Trie, result *[][]int) {
	for i := 0; i < len(word); i++ {
		if root.Index >= 0 && root.Index != index && isPalindrome(word, i, len(word)-1) {
			*result = append(*result, []int{index, root.Index})
		}
		root = root.Next[int(word[i]-'a')]
		// ba  Trie is a: [b] is from end to start
		// when it comes 'b' root is null
		if root == nil {
			return
		}
	}
	// whole string match, like, "tab" vs "bat"
	// "t" is not a word, "ta" is Not a word . "tab" is a word , but the rest of 'tab' is ""
	// and "" is Not palindome
	for i := range root.List {
		if root.List[i] != index {
			*result = append(*result, []int{index, root.List[i]})
		}
	}
}

func isPalindrome(word string, start, end int) bool {
	for start < end {
		if word[start] != word[end] {
			return false
		}
		start++
		end--
	}
	return true
}
