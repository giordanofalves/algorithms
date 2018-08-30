package main

import "fmt"

type sort []int

func newSorting() sort {
	numbers := []int{5, 2, 4, 6, 1, 3}

	return numbers
}

func (s sort) insertion() {
	for j := 2; j < len(s); j++ {
		key := s[j]
		i := j - 1

		for i > 0 && s[i] > key {
			s[i+1] = s[i]
			i = i - 1
		}

		s[i+1] = key
	}
}

func (s sort) print() {
	for _, number := range s {
		fmt.Println(number)
	}
}
