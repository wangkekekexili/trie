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

func (t *trie) Delete(k string) bool {
	runes := []rune(k)
	deleted, _ := t.root.delete(runes)
	return deleted
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

// delete returns two booleans.
// The first indicates whether k exists in the tree.
// The second indicates whether parent node should do cleanup.
func (n *node) delete(k []rune) (bool, bool) {
	if len(k) == 0 {
		if n.hasValue {
			n.hasValue = false
			n.value = nil
			return true, len(n.children) == 0
		} else {
			return false, false
		}
	}

	nextKey := k[0]
	nextNode, ok := n.children[nextKey]
	if !ok {
		return false, false
	}
	deleted, needCleanup := nextNode.delete(k[1:])
	if !needCleanup {
		return deleted, false
	}

	delete(n.children, nextKey)

	// The current node needs to be cleaned up if it doesn't have more children and doesn't have data.
	return true, len(n.children) == 0 && !n.hasValue
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
