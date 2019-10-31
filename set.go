package set

import (
	"fmt"
)

const maxSetSize = 10

//Set is our struct which represents a set data structure
type Set struct {
	Size   int
	Cursor int
	Items  [maxSetSize]string
}

//New instantiates an instance of a Set and returns it to the caller
func New() Set {
	return Set{
		Size:   0,
		Cursor: 0,
	}
}

//Add adds a string to a Set
func (s *Set) Add(t string) {
	if !(s.Contains(t)) {
		s.Items[s.Cursor] = t
		s.Cursor++
		if s.Cursor == maxSetSize {
			s.Cursor = 0
		}
		if s.Size < maxSetSize {
			s.Size++
		}
	}
}

//Remove removes string t from a Set if it exists
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
	return fmt.Errorf("set did not contain expected value %s after iterating, something is wrong", t)
}

//Contains checks to see if a Set already contains string t
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
