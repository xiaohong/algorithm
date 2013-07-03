package trie

type node struct{
	Key string
	Value interface{}
	childs []*node
}

type Trie struct{
	root *node
}

func NewTrie() *Trie{
	return &Trie{&node{"", nil, make(*node,0)}}
}