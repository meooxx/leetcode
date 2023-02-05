package main

import "fmt"

func main() {
	wd := Constructor()
	wd.AddWord("bad")

	fmt.Println(wd.Search(".ad"))
}
// 用a-z 替换 . 
// TL
type WordDictionary struct {
	mp       map[string]bool
	leadChar []map[byte]bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		mp: map[string]bool{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.mp[word] = true
	for i := range word {
		if len(this.leadChar) <= i {
			this.leadChar = append(this.leadChar, map[byte]bool{})
		}
		this.leadChar[i][word[i]] = true
	}
}

func (this *WordDictionary) Search(word string) bool {
	var searchImpl func(w string, start int) bool
	searchImpl = func(w string, start int) bool {
		if this.mp[w] {
			return true
		}
		for i := start; i < len(w); i++ {
			if i >= len(this.leadChar) {
				return false
			}
			if w[i] == '.' {
				for j := 'a'; j <= 'z'; j++ {
					by := []byte(w)
					by[i] = byte(j)
					if !this.leadChar[i][byte(j)] {
						continue
					}
					if searchImpl(string(by), i+1) {
						return true
					}
				}
			}
		}
		return false
	}
	return searchImpl(word, 0)
}

/**
* Your WordDictionary object will be instantiated and called as such:
* obj := Constructor();
* obj.AddWord(word);
* param_2 := obj.Search(word);
 */
