package main

func main() {
	findWords([][]byte{{'a', 'a'}}, []string{"aaa"})
}

func findWords(board [][]byte, words []string) []string {
	type Trie struct {
		word string
		sub  [26]*Trie
	}
	root := &Trie{}
	for i := range words {
		w := words[i]
		cur := root
		for j := range w {
			char := w[j] - 'a'
			if cur.sub[char] == nil {
				cur.sub[char] = &Trie{}
			}
			cur = cur.sub[char]
		}
		cur.word = w
	}
	result := []string{}
	var dfs func(board [][]byte, row, col int, t *Trie)
	dfs = func(board [][]byte, row, col int, t *Trie) {
		if row >= len(board) || row < 0 || col >= len(board[0]) || col < 0 {
			return
		}
		if board[row][col] == '#' {
			return
		}
		char := board[row][col] - 'a'
		if t.sub[char] == nil {
			return
		}
		if t.sub[char].word != "" {
			result = append(result, t.sub[char].word)
			t.sub[char].word = ""
		}
		board[row][col] = '#'
		dfs(board, row+1, col, t.sub[char])
		dfs(board, row-1, col, t.sub[char])
		dfs(board, row, col+1, t.sub[char])
		dfs(board, row, col-1, t.sub[char])
		board[row][col] = char + 'a'
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			dfs(board, row, col, root)
		}
	}
	return result
}
