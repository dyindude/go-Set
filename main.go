package main

import (
	"fmt"
)

const MAX_SIZE = 10

type Set struct {
	Size  int
	Items [MAX_SIZE]string
}

func New() Set {
	return Set{
		Size: 0,
	}
}

func (s *Set) Add(t string) {
	if !(s.Contains(t)) {
		s.Items[s.Size] = t
		s.Size++
	}
}

func (s *Set) Remove(t string) error {
	if !(s.Contains(t)) {
		return fmt.Errorf("set does not contain string %s", t)
	}
	for i, v := range s.Items {
		if v == t {
			s.Items[i] = ""
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
