package main

import "fmt"

func findWords1(board [][]byte, words []string) []string {
	result := []string{}
	var nextChar func(board [][]byte, word string, charIndex, row, col int) bool
	nextChar = func(board [][]byte, word string, charIndex, row, col int) bool {
		if charIndex >= len(word) {
			return true
		}
		if row >= len(board) || col >= len(board[0]) || row < 0 || col < 0 {
			return false
		}
		if board[row][col] != word[charIndex] {
			return false
		}
		board[row][col] = '#'
		result := nextChar(board, word, charIndex+1, row+1, col) || nextChar(board, word, charIndex+1, row-1, col) || nextChar(board, word, charIndex+1, row, col+1) || nextChar(board, word, charIndex+1, row, col-1)
		board[row][col] = '#'
		return result
	}
	for i := range words {
		word := words[i]
		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[row]); col++ {
				if board[row][col] == word[0] {
					if nextChar(board, word, i, row, col) {
						result = append(result, word)
					}
				}
			}

		}
	}
	return result

}

func main() {
	r := findWords([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"eat", "oath"})
	fmt.Println(r)
}
//   words      ad  be
//   root	 Trie [a, b, c]
//               d  e
//  board [a  b e]
//        [d  c h]
// loop board
// 1	[a] == Trie[0] => dfs向四个方向扫row+1[d] 命中 Trie[a].d 加入ad/
// 2	[b] == Tire[1] => ..... col+1 命中 Trie[b].e 加入be 
func findWords(board [][]byte, words []string) []string {
	type Trie struct {
		next [26]*Trie
		word string
	}
	result := []string{}
	root := &Trie{
		next: [26]*Trie{},
	}
	for i := range words {
		word := words[i]
		cur := root
		for j := range word {
			if cur.next[word[j]-'a'] == nil {
				cur.next[word[j]-'a'] = &Trie{
					next: [26]*Trie{},
				}
			}
			cur = cur.next[word[j]-'a']
		}
		cur.word = word
	}

	var dfs func(b [][]byte, row, col int, root *Trie)
	dfs = func(b [][]byte, row, col int, r *Trie) {
		if row < 0 || col < 0 || row >= len(b) || col >= len(b[0]) {
			return
		}
		char := b[row][col]
		by := char - 'a'
		
		if char == '#' || r.next[by] == nil {
			return
		}
		nextR := r.next[by]
		if nextR.word != "" {
			result = append(result, nextR.word)
			nextR.word = "" // de-duplicate
		}
		
		b[row][col] = '#'
		dfs(b, row+1, col, nextR)
		dfs(b, row-1, col, nextR)
		dfs(b, row, col+1, nextR)
		dfs(b, row, col-1, nextR)
		b[row][col] = char

	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			dfs(board, row, col, root)
		}
	}
	return result
}
