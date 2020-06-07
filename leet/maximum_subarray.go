package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Println("res:", res)

	nums = []int{-2, -1}
	res = maxSubArray(nums)
	fmt.Println("res:", res)

	// with index tracking
	fmt.Println("-=-=-=-")
	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum, start, end := maxSubArrayInd(nums)
	fmt.Println("sum:", sum, "start:", start, "end:", end)

	nums = []int{-2, -1}
	sum, start, end = maxSubArrayInd(nums)
	fmt.Println("sum:", sum, "start:", start, "end:", end)
}

// Kadane's algorithm
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxSoFar := nums[0]
	maxToHere := nums[0]

	for _, v := range nums {
		maxToHere += v
		maxToHere = int(math.Max(float64(maxToHere), float64(v)))

		maxSoFar = int(math.Max(float64(maxToHere), float64(maxSoFar)))
	}

	return maxSoFar
}

// keeping track of bounding indices
func maxSubArrayInd(nums []int) (int, int, int) {
	if len(nums) == 1 {
		return nums[0], 0, 0
	}

	maxSoFar := nums[0]
	maxSoFarStart := 0
	maxSoFarEnd := 0

	maxToHere := nums[0]
	maxToHereStart := nums[0]
	maxToHereEnd := nums[0]

	for i := 1; i < len(nums); i++ {
		maxToHere += nums[i]

		if maxToHere > nums[i] {
			maxToHereEnd++
		} else {
			maxToHere = nums[i]
			maxToHereStart = i
			maxToHereEnd = i
		}

		if maxToHere > maxSoFar {
			maxSoFar = maxToHere
			maxSoFarStart = maxToHereStart
			maxSoFarEnd = maxToHereEnd
		}
	}

	return maxSoFar, maxSoFarStart, maxSoFarEnd
}
