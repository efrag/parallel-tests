package search

func Iterative(nums []int, item int) int {
	low := 0
	high := len(nums) - 1

	for i := 0; i < 10000; i++ {
		continue
	}

	for low <= high {
		mid := low + ((high - low) / 2)

		if nums[mid] == item {
			return mid
		}

		if nums[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func IterativeWithError(nums []int, item int) int {
	if len(nums) == 3 {
		return -2
	}

	return Iterative(nums, item)
}
