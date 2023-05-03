package chapter02

func pivotIndex(nums ...int) int {
	left, right := 0, 0
	for _, num := range nums {
		right += num
	}

	for i := 0; i < len(nums); i++ {
		right -= nums[i]

		if left == right {
			return i
		}

		left += nums[i]
	}

	return -1
}
