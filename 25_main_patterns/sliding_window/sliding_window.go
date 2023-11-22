package sliding_window

import "reflect"

// This algorithmic technique is used when we need to handle the input data in specific window size.

// Task: given an array of integers and a positive integer k, find the maximum sum of a subarray of size k.
func maxSum(n []int, k int) int {
	var maxSum int

	for i := 0; i < len(n)-k+1; i++ {
		var currentSum int

		currentSlice := n[i : k+i]
		for _, value := range currentSlice {
			currentSum += value
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// Task: Given an array of positive integers and a target integer k, find the length of the smallest subarray with a sum greater than or equal to k. If there isn't one, return 0.
func minSum(n []int, k int) int {
	var minSum int

	for i := 0; i < len(n)-k+1; i++ {
		var currentSum int

		currenSlice := n[i : k+i]
		for _, value := range currenSlice {
			currentSum += value
		}

		if minSum == 0 || minSum > currentSum {
			minSum = currentSum
		} else if currentSum == minSum {
			return 0

		}
	}

	return minSum
}

// Task: Given a string and a pattern, find out if the string contains any permutation of the pattern.
func findPattern(s string, pattern string) bool {
	if len(s) < len(pattern) {
		return false
	}

	patternMap := make(map[byte]int)
	windowMap := make(map[byte]int)
	for i := 0; i < len(pattern); i++ {
		patternMap[pattern[i]]++
		windowMap[s[i]]++
	}

	for i := len(pattern); i < len(s); i++ {
		windowMap[s[i-len(pattern)]]--
		if windowMap[s[i-len(pattern)]] == 0 {
			delete(windowMap, s[i-len(pattern)])
		}
		windowMap[s[i]]++

		// to compare two maps we should use reflect.DeepEqual
		if reflect.DeepEqual(patternMap, windowMap) {
			return true
		}
	}

	return false
}
