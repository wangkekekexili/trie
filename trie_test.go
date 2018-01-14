package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	strs := map[string]string{"ann": "hello", "anna": "world", "anne": "ke"}

	trie := New()
	for k, v := range strs {
		trie.Add(k, v)
	}

	for input, expValue := range strs {
		gotValue, ok := trie.Get(input)
		if !ok {
			t.Fatalf("expect value for key %v to exist", input)
		}
		if expValue != gotValue {
			t.Fatalf("expect %v for key %v; got %v", expValue, input, gotValue)
		}
	}

	expNonExist := []string{"", "b", "an", "annb"}
	for _, input := range expNonExist {
		_, got := trie.Get(input)
		if got {
			t.Fatalf("expect %v not exist in the trie", input)
		}
	}
}
