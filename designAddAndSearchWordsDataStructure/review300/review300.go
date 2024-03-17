package main

type WordDictionary struct {
	isWord  bool
	wordMap [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		wordMap: [26]*WordDictionary{},
	}

}

func (this *WordDictionary) AddWord(word string) {
	for i := range word {
		index := int(word[i] - 'a')
		if this.wordMap[index] == nil {
			this.wordMap[index] = &WordDictionary{}
		}
		this = this.wordMap[index]

	}
	this.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.isWord
	}
	for i := range word {
		index := int(word[i] - 'a')
		if word[i] == '.' {
			for j := range this.wordMap {
				if this.wordMap[j] != nil {
					if this.wordMap[j].Search(word[i+1:]) {
						return true
					}
				}

			}
			return false
		} else if this.wordMap[index] != nil {
			return this.wordMap[index].Search(word[i+1:])
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
