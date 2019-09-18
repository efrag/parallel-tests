package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSortSimple(t *testing.T) {
	t.Parallel()
	nums := []int32{1, 2, 3, 45, 6, 7}
	expe := []int32{1, 2, 3, 6, 7, 45}

	BubbleSort(nums)

	if !reflect.DeepEqual(nums, expe) {
		t.Errorf("List not sorted as expected: %v vs %v", expe, nums)
	}
}

func TestBubbleSortSimpleEmpty(t *testing.T) {
	t.Parallel()
	nums := make([]int32, 0)
	expe := make([]int32, 0)

	BubbleSort(nums)

	if !reflect.DeepEqual(nums, expe) {
		t.Errorf("List not sorted as expected: %v vs %v", expe, nums)
	}
}
