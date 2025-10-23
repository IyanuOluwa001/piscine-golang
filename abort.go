package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	n := len(nums)

	// Bubble sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// Return the median (middle value)
	return nums[2]
}
