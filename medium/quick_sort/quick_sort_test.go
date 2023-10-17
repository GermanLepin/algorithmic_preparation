package quick_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_quickSort(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			nums           = []int{1, -34, 35, 4, -6}
			expectedResult = []int{-34, -6, 1, 4, 35}
		)

		quickSort(nums)
		r.Equal(nums, expectedResult)
	})
}
