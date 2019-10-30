package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	empty := New()
	fmt.Printf("%T\n", empty)
}

func TestAdd(t *testing.T) {
	empty := New()
	one := New()
	many := New()

	one.Add("foo")
	many.Add("foo")
	many.Add("bar")

	if empty.Size != 0 {
		t.Errorf("empty.Size = %d; want 0", empty.Size)
	}
	if one.Size != 1 {
		t.Errorf("one.Size = %d; want 1", one.Size)
	}
	if !many.Size >= 2 {
		t.Errorf("many.Size = %d; want >= 2", many.Size)
	}

}