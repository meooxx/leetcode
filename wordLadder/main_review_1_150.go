package main

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	wordMap := map[string]bool{}

	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return 0
	}
	delete(wordMap, beginWord)
	isNeigbor := func(a, b string) bool {
		dif := 0
		for i := range a {
			if a[i] != b[i] {
				dif++
			}
			if dif > 1 {
				return false
			}
		}
		return dif == 1
	}
	queue := []string{beginWord}
	anwser := 0

	for len(queue) > 0 {
		qSize := len(queue)
		anwser++
		for i := 0; i < qSize; i++ {
			lastWord := queue[0]
			queue = queue[1:]
			for nextWord := range wordMap {
				if !isNeigbor(lastWord, nextWord) {
					continue
				}
				if nextWord == endWord {
					anwser++
					return anwser
				}
				queue = append(queue, nextWord)
				delete(wordMap, nextWord)
			}
		}
	}
	return 0
}

//  beginWord 
//hit *it   [hit]			 
//    h*t   [hot]  *
//    hi*           
//hot *ot   [dot]   *                    
//    h*t
//    ho*       
//dot *ot
//    d*t
//    do*  [cog] *          
func ladderLength(beginWord string, endWord string, wordList []string) int {
	graph := map[string][]string{}
	for _, w := range wordList {
		for i := range w {
			byteW := []byte(w)
			byteW[i] = '*'
			stringPatter := string(byteW)
			if arr, ok := graph[stringPatter]; ok {
				graph[stringPatter] = append(arr, w)
			} else {
				graph[stringPatter] = []string{w}
			}
		}
	}
	anwser := 1
	queue := []string{beginWord}
	visited := map[string]bool{}
	for len(queue) > 0 {
		qSize := len(queue)
		anwser++
		for i := 0; i < qSize; i++ {
			preWord := queue[0]
			queue = queue[1:]
			for i := range preWord {
				byteW := []byte(preWord)
				byteW[i] = '*'
				strW := string(byteW)
				if arr, ok := graph[strW]; ok {
					for _, nextWord := range arr {
						if visited[nextWord] {
							continue
						}

						if nextWord == endWord {
							return anwser
						}

						queue = append(queue, nextWord)
						visited[nextWord] = true
					}
				}
			}
		}

	}
	return 0
}
