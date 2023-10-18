package main

type Trie struct {
	Next  [26]*Trie
	List  []int
	Index int
}

func palindromePairs(words []string) [][]int {
	t := &Trie{
		Index: -1,
		List:  []int{},
		Next:  [26]*Trie{},
	}
	for i := range words {
		addWord(words[i], i, t)
	}
	result := [][]int{}
	for i := range words {
		search(words[i], i, t, &result)
	}

	return result
}

func addWord(word string, index int, t *Trie) {
	for i := len(word) - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if t.Next[idx] == nil {
			t.Next[idx] = &Trie{
				List:  []int{},
				Next:  [26]*Trie{},
				Index: -1,
			}
		}
		if isPalindrome(word, 0, i) {
			t.List = append(t.List, index)
		}

		t = t.Next[idx]

	}
	t.List = append(t.List, index)
	t.Index = index
}

func search(word string, index int, t *Trie, result *[][]int) {
	for i := 0; i < len(word); i++ {
		if t.Index != -1 && t.Index != index && isPalindrome(word, i, len(word)-1) {
			// []int{index, t.Index},
			// the order should same as full match case 
			// bad case: [a, '']
			// 		start from 'a', t.Index == 1, got [1, 0]
			// 		goto the full match case, and also got [1, 0]
			*result = append(*result, []int{index, t.Index})
		}

		t = t.Next[int(word[i]-'a')]
		if t == nil {
			return
		}
	}
	// for whole string match
	// e.g tab vs bat. b is not word, ba is not word, and isPalindrome(word, 2,2) is false
	for i := range t.List {
		if index != t.List[i] {
			*result = append(*result, []int{index, t.List[i]})
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
