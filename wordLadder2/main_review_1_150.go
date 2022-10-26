package main

import "fmt"

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	allWords := map[string]bool{}
	if _, ok := allWords[endWord]; !ok {
		return [][]string{}
	}
	for _, word := range wordList {
		allWords[word] = true
	}
	delete(allWords, beginWord)
	isNeigbour := func(a, b string) bool {
		dif := 0
		for i := range a {
			if a[i] != b[i] {
				dif++
				if dif > 1 {
					return false
				}
			}
		}
		return dif == 1
	}
	queue := []string{beginWord}
	allNodes := [][]string{}
	fond := false
	for len(queue) > 0 && !fond {
		qSize := len(queue)
		levelNodes := make([]string, qSize)
		copy(levelNodes, queue)
		allNodes = append(allNodes, levelNodes)
		for i := 0; i < qSize; i++ {
			lastWord := queue[0]
			queue = queue[1:]
			for w := range allWords {
				if isNeigbour(lastWord, w) {
					if w == endWord {
						fond = true
						break
					}
					queue = append(queue, w)
					delete(allWords, w)
				}
			}
		}
	}

	result := [][]string{{endWord}}
	//  [d]    [ b ]
	//                 [a]
	//  [e]    [ c ]
	for i := len(allNodes) - 1; i >= 0; i-- {
		size := len(result)
		for _, word := range allNodes[i] {
			for i := 0; i < size; i++ {
				nextWords := result[i]
				if isNeigbour(word, nextWords[0]) {
					item := []string{word}
					item = append(item, nextWords...)
					result = append(result, item)
				}
			}
		}
		result = result[size:]
	}
	return result

}

func findLadders1(beginWord string, endWord string, wordList []string) [][]string {

	anwser := [][]string{}
	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return anwser
	}
	delete(wordSet, beginWord)
	levelNodes := [][]string{}
	queue := []string{beginWord}
	isNeighbor := func(word1, word2 string) bool {
		diff := 0
		for i := range word1 {
			if word1[i] != word2[i] {
				diff++
			}
			if diff > 1 {
				return false
			}
		}
		return diff == 1
	}
	fond := false
	for len(queue) > 0 && !fond {
		qSize := len(queue)
		nodes := make([]string, len(queue))
		copy(nodes, queue)
		levelNodes = append(levelNodes, nodes)

		for i := 0; i < qSize; i++ {
			curWord := queue[0]
			queue = queue[1:]

			for word := range wordSet {
				if !isNeighbor(curWord, word) {
					continue
				}
				if word == endWord {
					fond = true
					break
				}
				delete(wordSet, word)
				queue = append(queue, word)
			}
		}

	}
	anwser = [][]string{{endWord}}
	for i := len(levelNodes) - 1; i >= 0; i-- {
		size := len(anwser)
		for _, word := range levelNodes[i] {
			for j := 0; j < size; j++ {
				lastLevelWords := anwser[j]
				lastWord := lastLevelWords[0]
				if isNeighbor(word, lastWord) {
					item := []string{word}
					item = append(item, lastLevelWords...)
					anwser = append(anwser, item)
				}

			}
		}
		anwser = anwser[size:]

	}
	return anwser
}
