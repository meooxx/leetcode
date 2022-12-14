package main
/*
// a b c
// b a b

mp[a]				[a   	nil  nil]         insert abc
  [a][b] 		[nil, b, 	nil]
	[a][b][c] [nil, nil, c]
	
mp[b]				[a   		 	b   nil]        insert bab
	[b][a] 		[a, 			nil, nil]								 
	[b][a][c] [nil, nil, c]
*/
type Trie struct {
	mp     [26]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{
		mp:     [26]*Trie{},
		isWord: false,
	}
}

// { mp:[]}
func (this *Trie) Insert(word string) {
	for i := range word {
		if this.mp[word[i]-'a'] == nil {
			this.mp[word[i]-'a'] = &Trie{}
		}
		this = this.mp[word[i]-'a']
	}
	this.isWord = true
}

func (this *Trie) Search(word string) bool {
	for i := range word {
		if this.mp[word[i]-'a'] == nil {
			return false
		}
		this = this.mp[word[i]]
	}
	return this.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		if this.mp[prefix[i]] == nil {
			return false
		}
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
