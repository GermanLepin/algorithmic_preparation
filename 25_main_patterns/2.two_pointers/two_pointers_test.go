package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortedSquares(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			nums           = []int{-2, -1, 0, 2, 3}
			expectedResult = []int{0, 1, 4, 4, 9}
		)

		result := sortedSquares(nums)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			nums           = []int{-3, -1, 0, 1, 2}
			expectedResult = []int{0, 1, 1, 4, 9}
		)

		result := sortedSquares(nums)
		r.Equal(result, expectedResult)
	})

	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			nums           = []int{1, 1, 4, 9}
			expectedResult = []int{1, 1, 16, 81}
		)

		result := sortedSquares(nums)
		r.Equal(result, expectedResult)
	})
}

func Test_removeDuplicates(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			nums           = []int{-2, -1, -1, 2, 2, 2, 3, 3, 4, 4, 6, 7, 8, 8, 8, 9, 12}
			expectedResult = 10
		)

		result := removeDuplicates(nums)
		r.Equal(result, expectedResult)
	})
}
