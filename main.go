package main

import (
	"fmt"
)

const maxSetSize = 10

type Set struct {
	Size   int
	Cursor int
	Items  [maxSetSize]string
}

func New() Set {
	return Set{
		Size:   0,
		Cursor: 0,
	}
}

func (s *Set) Add(t string) {
	if !(s.Contains(t)) {
		if s.Size == maxSetSize {
			s.Items[s.Cursor] = t
			if s.Cursor == (maxSetSize - 1) {
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
			for n := i; n < (maxSetSize - 1); n++ {
				s.Items[n] = s.Items[n+1] //left shift
			}
			s.Items[maxSetSize-1] = ""
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
