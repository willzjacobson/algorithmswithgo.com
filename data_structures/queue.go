package main

import "fmt"

type Queue struct {
	array []int
	head  int
	tail  int
	size  int
}

func CreateQueue() *Queue {
	return &Queue{
		array: make([]int, 2),
	}
}

func (q *Queue) Enqueue(v int) {
	if q.isFull() {
		q.sizeUp()
	}

	lastInd := len(q.array) - 1
	if q.tail == lastInd && q.size > 0 {
		q.tail = 0
	} else if q.size > 0 {
		q.tail++
	}

	q.array[q.tail] = v
	q.size++
}

func (q *Queue) Dequeue() int {
	if q.size == 0 {
		panic("Queue underflow")
	}

	var el int
	lastInd := len(q.array) - 1
	if q.head == lastInd {
		q.head = 0
		el = q.array[lastInd]
	} else {
		q.head++
		el = q.array[q.head-1]
	}

	q.size--
	return el
}

func (q *Queue) isFull() bool {
	if q.head == q.tail+1 || (q.head == 0 && q.tail == len(q.array)-1) {
		return true
	}
	return false
}

func (q *Queue) sizeUp() {
	l := len(q.array)
	last := len(q.array) - 1
	newArr := make([]int, 2*l)

	for i := 0; i <= last; i++ {
		var src int
		if q.head+i <= last {
			src = q.array[q.head+i]
		} else {
			src = q.array[i+q.head-l]
		}
		newArr[i] = src
	}

	q.array = newArr
	q.head = 0
	q.tail = last
}

func (q *Queue) Size() int {
	return q.size
}

func main() {
	q := CreateQueue()
	fmt.Println("queue:", q)
	q.Enqueue(1)
	fmt.Println("queue:", q)
	q.Enqueue(2)
	fmt.Println("queue:", q)
	fmt.Println("poppin:", q.Dequeue())
	fmt.Println("queue:", q)
	q.Enqueue(3)
	fmt.Println("queue:", q)
	q.Enqueue(4)
	fmt.Println("queue:", q)
	fmt.Println("poppin:", q.Dequeue())
	fmt.Println("poppin:", q.Dequeue())
	fmt.Println("queue:", q)
	q.Enqueue(5)
	fmt.Println("queue:", q)
	q.Enqueue(6)
	fmt.Println("queue:", q)
	q.Enqueue(7)
	fmt.Println("queue:", q)
	q.Enqueue(8)
	fmt.Println("queue:", q)
	q.Enqueue(9)
	fmt.Println("queue:", q)
	q.Enqueue(10)
	fmt.Println("queue:", q)
}
