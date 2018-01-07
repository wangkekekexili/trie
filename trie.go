package trie

type trie struct {
	root *node
}

func New() *trie {
	return &trie{root: newNode()}
}

func (t *trie) Add(k string, v interface{}) {
	runes := []rune(k)
	t.root.add(runes, v)
}

func (t *trie) IsMember(k string) bool {
	runes := []rune(k)
	return t.root.isMember(runes)
}

type node struct {
	hasValue bool
	value    interface{}

	children map[rune]*node
}

func newNode() *node {
	return &node{
		children: make(map[rune]*node),
	}
}

func (n *node) add(k []rune, v interface{}) {
	if len(k) == 0 {
		n.hasValue = true
		n.value = v
		return
	}

	nextKey := k[0]
	if n.children[nextKey] == nil {
		n.children[nextKey] = newNode()
	}
	n.children[nextKey].add(k[1:], v)
}

func (n *node) isMember(k []rune) bool {
	if len(k) == 0 {
		return n.hasValue
	}
	nextNode, ok := n.children[k[0]]
	if !ok {
		return false
	}
	return nextNode.isMember(k[1:])
}
