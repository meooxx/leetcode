package main

import "fmt"

func main() {
	fmt.Println(ladderLength1("hit", "cog", []string{
		"hot", "dot", "dog", "lot", "log", "cog",
	}))
	fmt.Println(ladderLength1("a", "c", []string{
		"a", "b", "c",
	}))
}

// leet code for less runtime version
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	anwser := 0

	existed := false
	for _, word := range wordList {
		if word == endWord {
			existed = true
		}
	}
	if !existed {
		return anwser
	}

	graph := map[string][]string{}
	for _, word := range wordList {
		for i := range word {

			byteKey := []byte(word)
			byteKey[i] = '*'
			key := string(byteKey)
			if graph[key] != nil {
				graph[key] = append(graph[key], word)
			} else {
				graph[key] = []string{word}
			}
		}
	}
	queue := []string{beginWord}
	visited := map[string]bool{}
	for len(queue) > 0 {
		qSize := len(queue)
		anwser++
		for i := 0; i < qSize; i++ {
			lastWord := queue[0]
			queue = queue[1:]
			for i := range lastWord {
				byteKey := []byte(lastWord)
				byteKey[i] = '*'
				key := string(byteKey)
				if graph[key] != nil {
					for _, word := range graph[key] {
						if word == endWord {
							anwser++
							return anwser
						}
						if visited[word] {
							continue
						}
						queue = append(queue, word)
						visited[word] = true

					}
				}
			}
		}
	}
	return 0
}

// 从word ladder2 变化而来, 适应于wordladder2
func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	anwser := 0
	if !wordSet[endWord] {
		return anwser
	}

	delete(wordSet, beginWord)

	queue := []string{beginWord}

	isNeighbor := func(a, b string) bool {
		diff := 0
		for i := range a {
			if a[i] != b[i] {
				diff++
				if diff > 1 {
					return false
				}
			}
		}
		return diff == 1
	}

	for len(queue) > 0 {
		qSize := len(queue)
		anwser++
		for i := 0; i < qSize; i++ {
			lastWord := queue[i]
			for word := range wordSet {
				if word == endWord {
					anwser++
					return anwser
				}
				if !isNeighbor(word, lastWord) {
					continue
				}

				queue = append(queue, word)
				delete(wordSet, word)
			}
		}
		queue = queue[qSize:]
	}
	return 0
}
