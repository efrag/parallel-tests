package search

import "testing"

func TestTable(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		list     []int
		element  int
		position int
	}{
		{[]int{10, 20, 30, 40, 50, 60}, 30, 2},
		{[]int{10, 20, 30, 40, 50, 60}, 35, -1},
		{[]int{}, 65, -1},
	}

	for _, tc := range testCases {
		result := Iterative(tc.list, tc.element)
		if result != tc.position {
			msg := "Used list: %v " +
				"looked for: %d " +
				"position returned: %d"
			t.Errorf(msg, tc.list, tc.element, result)
		}
	}
}
