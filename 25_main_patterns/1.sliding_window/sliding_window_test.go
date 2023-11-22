package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSum(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{1, 3, 6, 4, 7, 4, 2, 6, 3, 7, 23, 4, 2, 6, 7, 34, 2, 6, 34}
			k              = 3
			expectedResult = 47
		)

		result := maxSum(num1, k)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{1, 3, 6, 4}
			k              = 2
			expectedResult = 10
		)

		result := maxSum(num1, k)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{-1, -3, 6, -4, 7, -2}
			k              = 3
			expectedResult = 9
		)

		result := maxSum(num1, k)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{41, 43, 36, 14, 47, 12}
			k              = 4
			expectedResult = 140
		)

		result := maxSum(num1, k)
		r.Equal(result, expectedResult)
	})
}

func Test_minSum(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{1, 3, 6, 4}
			k              = 2
			expectedResult = 4
		)

		result := minSum(num1, k)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{1, 3, 6, 4}
			k              = 3
			expectedResult = 10
		)

		result := minSum(num1, k)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{-1, -3, 6, -4, 7, -2}
			k              = 4
			expectedResult = -2
		)

		result := minSum(num1, k)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			num1           = []int{1, 3, 3, 1}
			k              = 2
			expectedResult = 0
		)

		result := minSum(num1, k)
		r.Equal(result, expectedResult)
	})
}

func Test_findPattern(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			s              = "oidbcaf"
			pattern        = "abc"
			expectedResult = true
		)

		result := findPattern(s, pattern)
		r.Equal(result, expectedResult)
	})

	t.Run("unsuccess test", func(t *testing.T) {
		r := require.New(t)

		var (
			s              = "oi"
			pattern        = "abc"
			expectedResult = false
		)

		result := findPattern(s, pattern)
		r.Equal(result, expectedResult)
	})

	t.Run("unsuccess test", func(t *testing.T) {
		r := require.New(t)

		var (
			s              = "oidbcaf"
			pattern        = "gfd"
			expectedResult = false
		)

		result := findPattern(s, pattern)
		r.Equal(result, expectedResult)
	})

	t.Run("unsuccess test", func(t *testing.T) {
		r := require.New(t)

		var (
			s              = "oidbcaf"
			pattern        = "gfd"
			expectedResult = false
		)

		result := findPattern(s, pattern)
		r.Equal(result, expectedResult)
	})

	t.Run("unsuccess test", func(t *testing.T) {
		r := require.New(t)

		var (
			s              = "aaabbbccccd"
			pattern        = "abc"
			expectedResult = false
		)

		result := findPattern(s, pattern)
		r.Equal(result, expectedResult)
	})
}
