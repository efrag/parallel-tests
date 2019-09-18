package sort

func BubbleSort(nums []int32) {
	for n := len(nums); n > 1; n-- {
		for j := 1; j < n; j++ {
			if nums[j-1] > nums[j] {
				swap(nums, j-1, j)
			}
		}
	}
}

func swap(nums []int32, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
