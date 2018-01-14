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

func (t *trie) Get(k string) (interface{}, bool) {
	runes := []rune(k)
	return t.root.get(runes)
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

func (n *node) get(k []rune) (interface{}, bool) {
	if len(k) == 0 {
		if !n.hasValue {
			return nil, false
		}
		return n.value, true
	}

	nextNode, ok := n.children[k[0]]
	if !ok {
		return nil, false
	}
	return nextNode.get(k[1:])
}
