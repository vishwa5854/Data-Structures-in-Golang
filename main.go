package main

import (
	"fmt"
	"os"

	binarysearchtree "heap-ds-go/m/binary-search-tree"
	hashtable "heap-ds-go/m/hash-table"
	"heap-ds-go/m/heap"
	linkedlist "heap-ds-go/m/linked-list"
	"heap-ds-go/m/queue"
	"heap-ds-go/m/stack"
	"heap-ds-go/m/trie"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage, only one option")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "linked-list":
		linkedlist.LinkedListDriver()
	case "heap":
		heap.HeapDriver()
	case "stack":
		stack.StackHandler()
	case "queue":
		queue.QueueHandler()
	case "bst":
		binarysearchtree.BinarySearchTreeHandler()
	case "trie":
		trie.TrieHandler()
	case "hash-table":
		hashtable.HashTableHandler()
	default:
		fmt.Println("Invalid Choice")
	}
}
