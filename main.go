package main

import (
	"fmt"
)

const MAX_SIZE = 10

type Set struct {
	Size   int
	Cursor int
	Items  [MAX_SIZE]string
}

func New() Set {
	return Set{
		Size:   0,
		Cursor: 0,
	}
}

func (s *Set) Add(t string) {
	if !(s.Contains(t)) {
		if s.Size == MAX_SIZE {
			s.Items[s.Cursor] = t
			if s.Cursor == (MAX_SIZE - 1) {
				s.Cursor = 0
			} else {
				s.Cursor++
			}
		} else {
			s.Items[s.Size] = t
			s.Size++
		}
	}
}

func (s *Set) Remove(t string) error {
	if !(s.Contains(t)) {
		return fmt.Errorf("set does not contain string %s", t)
	}
	for i, v := range s.Items {
		if v == t {
			for n := i; n < (MAX_SIZE - 1); n++ {
				s.Items[n] = s.Items[n+1] //left shift
			}
			s.Items[MAX_SIZE-1] = ""
			s.Size--
			return nil
		}
	}
	return fmt.Errorf("set did not contain expected value %s after iterating, something is wrong!", t)
}

func (s *Set) Contains(t string) bool {
	if s.Size == 0 {
		return false
	}
	for _, v := range s.Items {
		if v == t {
			return true
		}
	}
	return false
}
