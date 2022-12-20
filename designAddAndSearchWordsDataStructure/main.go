package main

import "fmt"

func main() {
	wd := Constructor()
	wd.AddWord("bad")

	fmt.Println(wd.Search(".ad"))
}

type WordDictionary struct {
	mp     [26]*WordDictionary
	isWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		mp:     [26]*WordDictionary{},
		isWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {

	for i := range word {
		index := word[i] - 'a'
		if this.mp[index] == nil {
			this.mp[index] = &WordDictionary{}
		}
		this = this.mp[index]
	}
	this.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	var searchImpl func(wd *WordDictionary, word string, i int) bool
	searchImpl = func(wd *WordDictionary, word string, i int) bool {
		if i == len(word) {
			return wd.isWord
		}
		if word[i] == '.' {
			for j := 0; j <= 25; j++ {
				w := wd
				if w.mp[j] != nil {
					w = w.mp[j]
					if searchImpl(w, word, i+1) {
						return true
					}
				}
			}
			return false
		}
		if wd.mp[word[i]-'a'] != nil {
			return searchImpl(wd.mp[word[i]-'a'], word, i+1)
		}

		return false
	}

	return searchImpl(this, word, 0)
}
