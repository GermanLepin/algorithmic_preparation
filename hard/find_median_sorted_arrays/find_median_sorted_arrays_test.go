package find_median_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMedianSortedArrays(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		var (
			nums1          = []int{1, 2, 3}
			nums2          = []int{1, 2, 5, 6}
			expectedResult = 3 // 1,2,3,5,6
		)

		result := findMedianSortedArrays(nums1, nums2)
		r.Equal(result, expectedResult)
	})
}
