package heap

import "fmt"

// MaxHeap struct which has a slice that holds the reference to the array
type MaxHeap struct {
	array []int
}

// Insert into heap
func (h *MaxHeap) Insert(key int) {
	/*
	 * 1. Add the element to the end of the array
	 * 2. Heapify to re-arrange the elements based on newly inserted element
	 */
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	if len(h.array) == 0 {
		fmt.Println("Cannot extract because the array length is 0")
		return -1
	}

	// heapify
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.maxHeapifyDown(0)

	return extracted
}

// heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	var leftChild, rightChild, childToCompare int = left(index), right(index), -1
	lastIndex := len(h.array) - 1

	for leftChild <= lastIndex {
		if leftChild == lastIndex { // when left child is the only child
			childToCompare = leftChild
		} else if h.array[leftChild] > h.array[rightChild] {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftChild, rightChild = left(index), right(index)
		} else {
			return
		}
	}
}

// Get parent index
func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// extract returns the largest key, and removes from the heap

func HeapDriver() {
	var arr = [3]int{10, 20, 30}
	m := &MaxHeap{}
	fmt.Println((m))

	for _, v := range arr {
		m.Insert(v)
		fmt.Println(m)
	}
}
