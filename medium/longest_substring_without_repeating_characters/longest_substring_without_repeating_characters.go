package longest_substring_without_repeating_characters

// Given a string s, find the length of the longest substring without repeating characters.

// leetcode https://leetcode.com/problems/longest-substring-without-repeating-characters/description

func lengthOfLongestSubstring(str string) int {
	letterMap := make(map[byte]int)
	result := 0

	for left, right := 0, 0; right < len(str); right++ {
		if index, ok := letterMap[str[right]]; ok {
			// only update the left pointer if it's behind the last position of the visited character
			left = max(left, index)
		}

		result = max(result, right-left+1)
		letterMap[str[right]] = right + 1
	}

	return result
}

func max(n, m int) int {
	if n > m {
		return n
	}

	return m
}
