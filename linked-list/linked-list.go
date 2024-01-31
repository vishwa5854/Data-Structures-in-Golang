package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

/** Reciever is going to be the LinkedList type */
func (l *LinkedList) prepend(n *node) {
	n.next = l.head
	l.head = n
	l.length += 1
}

func (l *LinkedList) append(n *node) {}

func (l LinkedList) printList() {
	fmt.Println("**********")
	for i := 0; i < l.length; i++ {
		fmt.Println(l.head.data)
		l.head = l.head.next
	}
	fmt.Println("Length of the list is ", l.length)
	fmt.Println("**********")
}

func (l *LinkedList) deleteWithValue(value int) {
	headPtr := l.head
	var prev *node = nil

	for i := 0; i < l.length; i++ {
		if headPtr.data == value {
			l.length -= 1

			if prev != nil {
				prev.next = headPtr.next
			} else {
				// this is the first element so we have to make second element as head
				l.head = headPtr.next
			}
			break
		} else {
			prev = headPtr
			headPtr = headPtr.next
		}
	}
}

func LinkedListDriver() {
	myList := &LinkedList{}
	node1 := &node{data: 1}

	myList.printList()

	myList.prepend(node1)
	myList.printList()

	myList.prepend(&node{data: 2})
	myList.printList()

	myList.prepend(&node{data: 3})

	myList.printList()

	myList.deleteWithValue(100)
	myList.printList()

	myList.deleteWithValue(3)
	myList.printList()

	myList.deleteWithValue(1)
	myList.printList()

	myList.deleteWithValue(2)
	myList.printList()

	myList.deleteWithValue(2)
	myList.printList()
}
