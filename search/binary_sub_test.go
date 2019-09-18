package search

import "testing"

func TestSubBasic(t *testing.T) {
	t.Parallel()
	in := []int{10, 20, 30, 40, 50, 60}

	t.Run("Item found", func(t *testing.T) {
		t.Parallel()
		result := Iterative(in, 30)
		if result != 2 {
			t.Error("wrong position")
		}
	})

	t.Run("Item not found", func(t *testing.T) {
		t.Parallel()
		result := Iterative(in, 35)
		if result != -1 {
			t.Error("item should not be found")
		}
	})
}

func TestSubDetails(t *testing.T) {
	t.Parallel()
	in := []int{10, 20, 30, 40, 50, 60}
	t.Log("Before the sub tests start")
	t.Run("Item found", func(t *testing.T) {
		t.Parallel()
		result := Iterative(in, 30)
		if result != 2 {
			t.Error("wrong position")
		}
	})

	t.Run("Item not found", func(t *testing.T) {
		t.Parallel()
		result := Iterative(in, 35)
		if result != -1 {
			t.Error("item should not be found")
		}
	})
	t.Log("After the sub tests finish")
}

func TestSubTable(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		list     []int
		element  int
		position int
	}{
		{"Found", []int{10, 20, 30, 40, 50, 60}, 30, 2},
		{"Not found", []int{10, 20, 30, 40, 50, 60}, 35, -1},
		{"Empty list", []int{}, 65, -1},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Iterative(tc.list, tc.element)
			if result != tc.position {
				msg := "List: %v item: %d position returned: %d"
				t.Errorf(msg, tc.list, tc.element, result)
			}
		})
	}
}
