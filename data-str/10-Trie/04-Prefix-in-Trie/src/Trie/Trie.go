package Trie

type node struct {
	isWord bool
	value  int
	next   map[string]*node
}

type Trie struct {
	root *node
	size int
}

func getNode() *node {
	return &node{next: make(map[string]*node)}
}

func Constructor() *Trie {
	return &Trie{root: getNode()}
}

// 获得Trie中存储的单词数量
func (this *Trie) GetSize() int {
	return this.size
}

// 向Trie中添加一个新的单词word
func (this *Trie) Add(word string, value int) {
	cur := this.root

	for _, w := range []rune(word) {
		c := string(w)

		if cur.next[c] == nil {
			cur.next[c] = getNode()
		}
		cur = cur.next[c]
	}

	if cur.isWord == false {
		cur.isWord = true
		cur.value = value
		this.size++
	}
}

// 查询单词word是否在Trie中
func (this *Trie) Contains(word string) bool {
	cur := this.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}

	return cur.isWord
}

// 查询是否在Trie中有单词以prefix为前缀
func (this *Trie) IsPrefix(prefix string) bool {
	cur := this.root

	for _, s := range []rune(prefix) {
		c := string(s)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}

	return true
}

/** Returns if the word is in the data structure. A word could contain the dot
character '.' to represent any one letter. */
func (this *Trie) Search(word string) bool {
	return this.match(this.root, word, 0)
}

func (this *Trie) match(n *node, word string, index int) bool {
	if index == len(word) {
		return n.isWord
	}
	c := string([]rune(word)[index])
	if c != "." {
		if n.next[c] == nil {
			return false
		}
		return this.match(n.next[c], word, index+1)
	} else {
		for w := range n.next {
			if this.match(n.next[w], word, index+1) {
				return true
			}
		}
		return false
	}
}

func (this *Trie) Sum(prefix string) int {
	cur := this.root
	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return 0
		}
		cur = cur.next[c]
	}

	return this.sum(cur)
}

func (this *Trie) sum(n *node) int {
	res := n.value
	for s := range n.next {
		res += this.sum(n.next[s])
	}
	return res
}
