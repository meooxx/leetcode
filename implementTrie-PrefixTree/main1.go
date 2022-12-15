package main

type Trie struct {
	mp map[string]bool
}

func Constructor() Trie {
	return Trie{
		mp: map[string]bool{},
	}
}

func (this *Trie) Insert(word string) {
	this.mp[word] = true
}

func (this *Trie) Search(word string) bool {
	return this.mp[word]
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := range this.mp {
		if strings.Index(i, prefix) == 0 {
			return true
		}
	}
	return false
}

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
 */