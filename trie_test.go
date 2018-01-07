package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := New()
	trie.Add("ann", 1)
	trie.Add("anna", 2)
	trie.Add("anne", 3)

	expExist := []string{"ann", "anna", "anne"}
	for _, input := range expExist {
		got := trie.IsMember(input)
		if !got {
			t.Fatalf("expect %v exist in the trie", input)
		}
	}

	expNonExist := []string{"", "b", "an", "annb"}
	for _, input := range expNonExist {
		got := trie.IsMember(input)
		if got {
			t.Fatalf("expect %v not exist in the trie", input)
		}
	}
}
