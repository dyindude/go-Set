package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	empty := New()
	fmt.Printf("%T\n", empty)
}

func ExampleData() (Set, Set, Set) {
	empty := New()
	one := New()
	many := New()

	one.Add("foo")
	many.Add("foo")
	many.Add("bar")

	return empty, one, many
}

func FullData() Set {
	full := New()
	full.Add("fee")
	full.Add("fie")
	full.Add("foo")
	full.Add("fum")
	full.Add("fuz")
	full.Add("bee")
	full.Add("bie")
	full.Add("boo")
	full.Add("bum")
	full.Add("buz")
	return full
}

func TestAdd(t *testing.T) {
	empty, one, many := ExampleData()

	if empty.Size != 0 {
		t.Errorf("empty.Size = %d; want 0", empty.Size)
	}
	if one.Size != 1 {
		t.Errorf("one.Size = %d; want 1", one.Size)
	}
	if !(many.Size >= 2) {
		t.Errorf("many.Size = %d; want >= 2", many.Size)
	}
}

func TestContains(t *testing.T) {
	empty, one, many := ExampleData()

	if empty.Contains("") {
		t.Errorf("empty contains an empty string \"\" which is not valid for our example.")
	}
	if !one.Contains("foo") {
		t.Errorf("one does not contain foo")
	}
	if !many.Contains("foo") || !many.Contains("bar") {
		t.Errorf("many does not contain foo or bar")
	}
}

func TestContainsNoDuplicates(t *testing.T) {
	_, one, many := ExampleData()

	one.Add("foo")
	many.Add("bar")
	found := 0
	for _, v := range one.Items {
		if v == "foo" {
			found++
		}
	}
	if found > 1 {
		t.Errorf("one contains %d instances of \"foo\", should only have 1", found)
	}
	found = 0
	for _, v := range many.Items {
		if v == "bar" {
			found++
		}
	}
	if found > 1 {
		t.Errorf("many contains %d instances of \"bar\", should only have 1", found)
	}
}

func TestRemove(t *testing.T) {
	_, one, many := ExampleData()

	err := one.Remove("foo")
	if err != nil {
		t.Errorf("error removing \"foo\" from one: %s", err)
	}
	if one.Size != 0 {
		t.Errorf("After removing \"foo\", one should be size 0. got %d", one.Size)
	}
	err = one.Remove("bar")
	if err == nil {
		t.Errorf("Remove() did not provide expected error when removing item not in Set")
	}
	err = many.Remove("foo")
	if err != nil {
		t.Errorf("error removing \"foo\" from many: %s", err)
	}
	if many.Size != 1 {
		t.Errorf("After removing \"foo\", many should be size 1. got %d", many.Size)
	}
	err = many.Remove("bar")
	if err != nil {
		t.Errorf("error removing \"bar\" from many: %s", err)
	}
	if many.Size != 0 {
		t.Errorf("After removing \"bar\", many should be size 0. got %d", many.Size)
	}
	err = many.Remove("baz")
	if err == nil {
		t.Errorf("Remove() did not provide expected error when removing item not in Set")
	}
}

func TestOverwriteOldest(t *testing.T) {
	full := FullData()
	full.Add("newvalue")
	if full.Size > MAX_SIZE {
		t.Errorf("full has an invalid size %d", full.Size)
	}
	if !(full.Contains("newvalue")) {
		t.Errorf("full does not contain \"newvalue\" after adding it.")
	}
	if full.Contains("fee") {
		t.Errorf("full still contains \"fee\" (the first value added) after adding another set item")
	}
}
