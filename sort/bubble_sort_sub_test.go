package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSortSub(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name         string
		unsortedList []int32
		sortedList   []int32
	}{
		{"Already sorted", []int32{1, 2, 3}, []int32{1, 2, 3}},
		{"Unsorted", []int32{2, 3, 1}, []int32{1, 2, 3}},
		{"Empty list", make([]int32, 0), make([]int32, 0)},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			BubbleSort(tc.unsortedList)
			if !reflect.DeepEqual(tc.unsortedList, tc.sortedList) {
				msg := "Expected list: %v got list: %v"
				t.Errorf(msg, tc.sortedList, tc.unsortedList)
			}
		})
	}
}
