package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	empty := New()
	fmt.Printf("%T\n", empty)
}
