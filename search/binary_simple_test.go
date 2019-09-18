package search

import "testing"

func TestSimpleFound(t *testing.T) {
	t.Parallel()
	in := []int{10, 20, 30, 40, 50, 60}

	result := Iterative(in, 30)
	if result != 2 {
		t.Error("wrong position")
	}
}

func TestSimpleNotFoundInRange(t *testing.T) {
	t.Parallel()
	in := []int{10, 20, 30, 40, 50, 60}

	result := Iterative(in, 35)
	if result != -1 {
		t.Error("item should not be found")
	}
}

func TestSimple(t *testing.T) {
	t.Parallel()
	in := []int{10, 20, 30, 40, 50, 60}

	result := Iterative(in, 30)
	if result != 2 {
		t.Error("wrong position")
	}

	result = Iterative(in, 35)
	if result != -1 {
		t.Error("item should not be found")
	}

	result = Iterative(in, 65)
	if result != -1 {
		t.Errorf("item %d "+
			"found at position %d "+
			"but does not exist in %v", 65, result, in)
	}
}
