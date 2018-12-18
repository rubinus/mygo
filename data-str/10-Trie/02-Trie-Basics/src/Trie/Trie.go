package Trie

type node struct {
	isWord bool
	next   map[string]*node
}

type Trie struct {
	root *node
	size int
}

func getNode() *node {
	next := make(map[string]*node)
	return &node{
		next: next,
	}
}

func Constructor() *Trie {
	return &Trie{
		root: getNode(),
	}
}

// 获得Trie中存储的单词数量
func (this *Trie) GetSize() int {
	return this.size
}

// 向Trie中添加一个新的单词word
func (this *Trie) Add(word string) {
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
		this.size++
	}
}
