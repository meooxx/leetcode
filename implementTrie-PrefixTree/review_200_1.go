package main

type Trie struct {
	sub    [26]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	// hello
	// h
	//  e
	//    l
	//      l
	//        o
	for i := range word {
		char := word[i] - 'a'
		if this.sub[char] == nil {
			this.sub[char] = &Trie{}
		}
		this = this.sub[char]
	}
	this.isWord = true
}

func (this *Trie) Search(word string) bool {
	for i := range word {
		char := word[i] - 'a'
		if this.sub[char] == nil {
			return false
		}
		this = this.sub[char]
	}
	return this.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		char := prefix[i] - 'a'
		if this.sub[char] == nil {
			return false
		}
		this = this.sub[char]
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
