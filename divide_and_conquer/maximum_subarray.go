package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum1, start1, end1 := maxSubArray(nums, 0, len(nums))
	fmt.Println("sum:", sum1, "start:", start1, "end:", end1)

	fmt.Println("-=-=-=-")

	nums1 := []int{-2, -1}
	sum2, start2, end2 := maxSubArray(nums1, 0, len(nums1))
	fmt.Println("sum:", sum2, "start:", start2, "end:", end2)
}

func maxSubArray(nums []int, start int, end int) (int, int, int) {
	l := len(nums)
	if l == 1 {
		return nums[0], start, end
	}

	mid := l / 2
	left := nums[:mid]
	right := nums[mid:]
	sumLeft, startLeft, endLeft := maxSubArray(left, 0, mid-1)
	sumRight, startRight, endRight := maxSubArray(right, mid, l-1)
	sumCross, startCross, endCross := maxCrossMidArray(nums, mid)

	if sumLeft > sumRight && sumLeft > sumCross {
		return sumLeft, startLeft, endLeft
	} else if sumRight > sumLeft && sumRight > sumCross {
		return sumRight, startRight, endRight
	} else {
		return sumCross, startCross, endCross
	}
}

func maxCrossMidArray(nums []int, mid int) (int, int, int) {
	leftSum := 0
	leftSumMax := 0
	leftStart := mid - 1
	for i := mid; i >= 0; i-- {
		leftSum += nums[i]
		if leftSum > leftSumMax {
			leftSumMax = leftSum
			leftStart = i
		}
	}

	rightSum := 0
	rightSumMax := 0
	rightEnd := mid
	for i := mid + 1; i < len(nums); i++ {
		rightSum += nums[i]
		if rightSum > rightSumMax {
			rightSumMax = rightSum
			rightEnd = i
		}
	}

	return leftSum + rightSum, leftStart, rightEnd
}
