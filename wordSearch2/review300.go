package main

type Trie struct {
	word string
	mp   [26]*Trie
}

func (this *Trie) Add(word string) {

	for i := range word {
		if this.mp[int(word[i]-'a')] == nil {
			this.mp[int(word[i]-'a')] = &Trie{}
		}
		this = this.mp[int(word[i]-'a')]
	}
	this.word = word
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for i := range words {
		root.Add(words[i])
	}
	result := []string{}
	var dfs func(board [][]byte, row, col int, root *Trie)
	dfs = func(board [][]byte, row, col int, root *Trie) {
		if row > len(board) || row < 0 || col > len(board[0]) || col < 0 {
			return
		}
		b := board[row][col] - 'a'
		if board[row][col] == '#' || root.mp[int(b)] == nil {
			return
		}
		next := root.mp[int(b)]
		if next.word != "" {
			result = append(result, next.word)
		}
		board[row][col] = '#'
		dfs(board, row, col+1, next)
		dfs(board, row, col-1, next)
		dfs(board, row+1, col, next)
		dfs(board, row-1, col, next)

	}
	for row := range board {
		for col := range board[row] {
			dfs(board, row, col, root)
		}
	}
	return result

}
