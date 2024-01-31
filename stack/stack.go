package stack

import "fmt"

type Stack struct {
	items []int
}

// push - add an element to back
func (s *Stack) push(data int) {
	s.items = append(s.items, data)
}

// pop - will remove an element from back & returns removed val
func (s *Stack) pop() int {
	stackLength := len(s.items) - 1

	if stackLength < 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	removedItem := s.items[stackLength]
	s.items = s.items[:stackLength]
	return removedItem
}

func StackHandler() {
	s := &Stack{}

	fmt.Println(s)
	s.push(1)
	fmt.Println(s)
	s.push(2)
	fmt.Println(s)
	s.push(3)
	fmt.Println(s)

	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())

}
