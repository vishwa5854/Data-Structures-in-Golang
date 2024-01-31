package binarysearchtree

import (
	"fmt"
)

/**
 * TODO
 * 1. What is Balanced Tree?
 * 2. What is Balanced Binary Tree?
 * 3. What are AVL trees?
 * 4. What is complete tree or complete binary tree?
 * 5. Once refer to B+ trees as well.
 */

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// insert
func (root *Node) insert(data int) {
	if data > root.Key {
		// right
		if root.Right == nil {
			root.Right = &Node{Key: data}
		} else {
			// right side is not empty
			root.Right.insert(data)
		}
	} else if data < root.Key {
		// left
		if root.Left == nil {
			root.Left = &Node{Key: data}
		} else {
			// left side is not empty
			root.Left.insert(data)
		}
	}
}

// search
func (root *Node) search(target int) bool {
	if root == nil {
		return false
	}
	if root.Key == target {
		return true
	} else if root.Key > target {
		return root.Left.search(target)
	} else if root.Key < target {
		return root.Right.search(target)
	}
	return false
}

// inorder recursive traversal
func (root *Node) rInOrder() {
	if root == nil {
		return
	}
	root.Left.rInOrder()
	fmt.Printf("%d, ", root.Key)
	root.Right.rInOrder()
}

// inorder iterative traversal using stack
// func (root *Node) iInOrder() {
// 	stack := &stack.Stack{}

// }

// preOrder recursive traversal
func (root *Node) rPreOrder() {
	if root == nil {
		return
	}
	fmt.Printf("%d, ", root.Key)
	root.Left.rPreOrder()
	root.Right.rPreOrder()
}

// postOrder recursive traversal
func (root *Node) rPostOrder() {
	if root == nil {
		return
	}
	root.Left.rPostOrder()
	root.Right.rPostOrder()
	fmt.Printf("%d, ", root.Key)
}

func BinarySearchTreeHandler() {
	tree := &Node{Key: 100}

	tree.insert(150)
	tree.insert(50)
	tree.insert(65)
	tree.insert(144)
	tree.insert(400)
	tree.insert(4)

	tree.rInOrder()
	fmt.Println()
	tree.rPreOrder()
	fmt.Println()
	tree.rPostOrder()
	fmt.Println()

	fmt.Println(tree.search(100))
	fmt.Println(tree.search(200))
	fmt.Println(tree.search(4))
}
