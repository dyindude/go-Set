package main

type Set struct {
	Size int
}

func New() Set {
	return Set{
		Size: 0,
	}
}
