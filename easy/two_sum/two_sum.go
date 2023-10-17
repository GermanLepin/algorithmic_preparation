package two_sum

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Leetcode https://leetcode.com/problems/add-two-numbers/

func TwoSumMapValues(nums []int, target int) []int {
	hMap := make(map[int]int)

	//O(n) --> hash map
	for i, num := range nums {
		_, ok := hMap[num]
		if ok {
			return []int{hMap[num], i}
		}

		hMap[target-num] = i
	}

	return nil
}
