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
	if len(nums) <= 1 {
		return len(nums)
	}

	uniqPosition := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[uniqPosition] {
			uniqPosition++
			nums[uniqPosition] = nums[i]
		}
	}

	return uniqPosition + 1
}

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which
// are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted,
// you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are
// not equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.

// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3]

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
