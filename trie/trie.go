package trie

import "fmt"

// Node
type Node struct {
	children     [26]*Node
	endOfTheWord bool
}

// Trie
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// insert
func (t *Trie) Insert(w string) {
	curr := t.root

	for _, v := range w {
		currLetterIndex := v - 'a'

		if curr.children[currLetterIndex] == nil {
			curr.children[currLetterIndex] = &Node{}
		}
		curr = curr.children[currLetterIndex]
	}
	curr.endOfTheWord = true
}

// search
func (t *Trie) Search(w string) bool {
	curr := t.root

	for _, v := range w {
		currLetterIndex := v - 'a'

		if curr.children[currLetterIndex] == nil {
			return false
		}
		curr = curr.children[currLetterIndex]
	}
	return curr.endOfTheWord
}

func TrieHandler() {
	tr := InitTrie()
	// fmt.Println(tr, tr.root)

	for _, v := range []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	} {
		tr.Insert(v)
	}
	fmt.Println(tr.Search("aragorn"))
	fmt.Println(tr.Search("aragorns"))
	fmt.Println(tr.Search("argon"))
	fmt.Println(tr.Search(""))
}
