package main

import "fmt"

type sort []int

func newSorting() sort {
	numbers := []int{6, 10, 4, 9, 1, 3, 2, 8, 5, 7}

	return numbers
}

func (s sort) insertion() {
	for j := 1; j < len(s); j++ {
		key := s[j]
		i := j - 1

		for i > -1 && s[i] > key {
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
