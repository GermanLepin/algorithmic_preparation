package sliding_window

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
// Permutation is defined as the re-arranging of the characters of the string.
// Input: String="oidbcaf", Pattern="ABC"
// Output: true
// Explanation: The string contains "bca" which is a permutation of the given pattern.
