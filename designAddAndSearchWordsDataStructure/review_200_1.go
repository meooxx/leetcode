package main

type WordDictionary struct {
	isWord bool
	sub    [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	for i := range word {
		char := word[i] - 'a'
		if this.sub[char] == nil {
			this.sub[char] = &WordDictionary{}
		}
		this = this.sub[char]
	}
	this.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return SearchImpl(this, word, 0)
}

func SearchImpl(wd *WordDictionary, word string, start int) bool {
	if wd == nil {
		return false
	}
	if start >= len(word) {
		return wd.isWord
	}
	for i := start; i < len(word); i++ {
		char := word[i] - 'a'
		if word[i] == '.' {
			// w may equal nil
			for _, w := range wd.sub {
				if SearchImpl(w, word, i+1) {
					return true
				}
			}
			return false
		} else if wd.sub[char] != nil {
			return SearchImpl(wd.sub[char], word, i+1)
		} else {
			return false
		}

	}
	return false
}

/**
* Your WordDictionary object will be instantiated and called as such:
* obj := Constructor();
* obj.AddWord(word);
* param_2 := obj.Search(word);
 */
