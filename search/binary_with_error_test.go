package search

import "testing"

func TestErrorSubTable(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		list     []int
		element  int
		position int
	}{
		{"Found", []int{10, 20, 30, 40, 50, 60}, 30, 2},
		{"Empty list", []int{}, 65, -1},
		{"3 items", []int{10, 20, 30}, 30, 2},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := IterativeWithError(tc.list, tc.element)
			if result != tc.position {
				msg := "List: %v item: %d position returned: %d"
				t.Errorf(msg, tc.list, tc.element, result)
			}
		})
	}
}
