package main

import "fmt"

type Stack struct {
	top   int
	array []int
}

func CreateStack() *Stack {
	return &Stack{
		array: []int{},
		top:   -1,
	}
}
func (s *Stack) Push(v int) {
	s.top++
	if len(s.array) > s.top {
		s.array[s.top] = v
	} else {
		s.array = append(s.array, v)
	}
}
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack underflow")
	}
	s.top--
	return s.array[s.top+1]
}
func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func main() {
	s := CreateStack()
	fmt.Println("stack:", s)
	fmt.Println("Is empty:", s.IsEmpty())
	s.Push(2)
	fmt.Println("stack:", s)
	fmt.Println("Is empty:", s.IsEmpty())
	s.Push(3)
	fmt.Println("stack:", s)
	fmt.Println("poppin:", s.Pop())
	fmt.Println("stack:", s)
	s.Push(4)
	fmt.Println("stack:", s)
	s.Push(5)
	fmt.Println("stack:", s)
}
