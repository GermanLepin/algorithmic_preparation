package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			example1       = "abcabcbb"
			expectedResult = 3 // "abc"
		)

		result := lengthOfLongestSubstring(example1)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			example1       = "asdfasdasa"
			expectedResult = 4 // "abc"
		)

		result := lengthOfLongestSubstring(example1)
		r.Equal(result, expectedResult)
	})
}
