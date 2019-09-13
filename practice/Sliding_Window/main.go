package main

import (
	"fmt"

	stack "stack.com/lib"
)

func printListMax(a []int, k int) {
	var j int = 0
	maxupto := make([]int, len(a))
	s := stack.New()
	s.Push(0)
	for i := 1; i < len(a); i++ {
		for a[i] > a[s.Peek()] {
			maxupto[s.Peek()] = i - 1
			s.Pop()
		}
		s.Push(i)
	}

	for s.Len() > 0 {
		maxupto[s.Peek()] = len(a) - 1
		s.Pop()
	}

	fmt.Println(maxupto)

	for i := 0; i <= len(a)-k; i++ {
		if j < i {
			j = i
		}
		for maxupto[j] < i+k-1 {
			j = maxupto[j] + 1
		}
		fmt.Println(a[j])
	}
}

func main() {
	fmt.Println("Hello World!")
	a := []int{9, 7, 2, 4, 6, 8, 2, 1, 5}
	k := 3
	printListMax(a, k)
}
