package two_pointers

// Usage: In this technique, we use two pointers to iterate the input data.
// Generally, both pointers move in the opposite direction at a constant interval.
// It works only with sorted data

// Task: Given a sorted array, create a new array containing squares of all the
// numbers of the input array in the sorted order.

// Example:
// Input: [-2, -1, 0, 2, 3]
// Output: [0, 1, 4, 4, 9]

func sortedSquares(nums []int) []int {
	sortedNums := make([]int, len(nums))

	var (
		left               = 0
		right              = len(nums) - 1
		highestSquareIndex = len(nums) - 1
	)

	for left <= right {
		leftSquard := nums[left] * nums[left]
		rightSquard := nums[right] * nums[right]

		if leftSquard > rightSquard {
			sortedNums[highestSquareIndex] = leftSquard
			left += 1
		} else {
			sortedNums[highestSquareIndex] = rightSquard
			right += -1
		}

		highestSquareIndex -= 1
	}

	return sortedNums
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Task: Given an integer array nums sorted in non-decreasing order, remove the duplicates
// in-place such that each unique element appears only once. The relative order of the
// elements should be kept the same. Then return the number of unique elements in nums.

func removeDuplicates(nums []int) int {

	return 0
}
