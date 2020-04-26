package main

type Trie struct {
	isWord bool
	Next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	tmp := this
	for i := 0; i < len(word); i++ {
		pos := word[i] - 'a'
		if tmp.Next[pos] == nil {
			tmp.Next[pos] = new(Trie)
		}
		tmp = tmp.Next[pos]
	}
	tmp.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	pos := 0
	tmp := this
	for pos < len(word) && tmp.Next[word[pos] - 'a'] != nil {
		tmp = tmp.Next[word[pos] - 'a']
		pos ++
	}
	if pos == len(word) && tmp.isWord {
		return true
	} else {
		return false
	}
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	pos := 0
	tmp := this
	for pos < len(prefix) && tmp.Next[prefix[pos] - 'a'] != nil {
		tmp = tmp.Next[prefix[pos] - 'a']
		pos ++
	}
	if pos == len(prefix) {
		return true
	} else {
		return false
	}
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
