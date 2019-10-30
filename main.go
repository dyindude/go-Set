package main

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
