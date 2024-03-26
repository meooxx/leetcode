package main

type Trie struct {
	mp     [26]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for i := range word {
		index := int(word[i] - 'a')
		if this.mp[index] == nil {
			this.mp[index] = &Trie{}
		}
		this = this.mp[index]
	}
	this.isWord = true
}

func (this *Trie) Search(word string) bool {
	for i := range word {
		index := int(word[i] - 'a')
		if this.mp[index] == nil {
			return false
		}
		this = this.mp[index]
	}
	return this.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		index := int(prefix[i] - 'a')
		if this.mp[index] == nil {
			return false
		}
		this = this.mp[index]
	}
	return true
}

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
 */
