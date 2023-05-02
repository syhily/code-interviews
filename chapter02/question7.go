package chapter02

import (
	"crypto/rand"
	"math/big"

	"github.com/syhily/code-interviews/common"
)

func threeSum(nums ...int) [][]int {
	quickSortArray(nums)
	var res [][]int

	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		r := innerThreeSum(nums, i)
		if len(r) == 3 {
			res = append(res, r)
		}
	}

	return res
}

func innerThreeSum(nums []int, index int) []int {
	target, s, e := -nums[index], index+1, len(nums)-1

	for s < e {
		sum := nums[s] + nums[e]
		if sum < target {
			s++
		} else if sum > target {
			e--
		} else {
			return []int{nums[index], nums[s], nums[e]}
		}
	}

	return nil
}

func quickSortArray(nums []int) {
	stack := common.NewStack[int]()
	stack.Push(0)
	stack.Push(len(nums) - 1)

	for !stack.Empty() {
		right := stack.Pop()
		left := stack.Pop()

		if left < right {
			middle := partitionArray(nums, left, right)
			stack.Push(left)
			stack.Push(middle - 1)
			stack.Push(middle + 1)
			stack.Push(right)
		}
	}
}

func partitionArray(nums []int, start, end int) int {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(end-start)))
	middle := int(r.Int64()) + start
	nums[start], nums[middle] = nums[middle], nums[start]

	middle = start
	for i := start + 1; i <= end; i++ {
		if nums[i] <= nums[middle] {
			middle++
			nums[middle], nums[i] = nums[i], nums[middle]
			nums[middle-1], nums[middle] = nums[middle], nums[middle-1]
		}
	}

	return middle
}
