package queue

import "fmt"

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// enqueue
func (q *Queue) enqueue(data int) {
	q.items = append(q.items, data)
}

// deque
func (q *Queue) deque() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	ans := q.items[0]
	q.items = q.items[1:]
	return ans
}

func QueueHandler() {
	fmt.Println("Queue handler")
	q := &Queue{}

	fmt.Println(q)

	q.enqueue(1)
	fmt.Println(q)

	q.enqueue(2)
	fmt.Println(q)

	q.enqueue(3)
	fmt.Println(q)

	fmt.Println(q.deque(), q)
	fmt.Println(q.deque(), q)
	fmt.Println(q.deque(), q)
	fmt.Println(q.deque(), q)
}
