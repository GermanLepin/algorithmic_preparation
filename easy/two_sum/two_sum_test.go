package two_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	t.Run("test two sum", func(t *testing.T) {
		r := require.New(t)

		expectedResult := []int{0, 4}
		nums := []int{2, 3, 1, 11, 4, 6, 7, 24, 9}
		target := 6

		result := TwoSumMapValues(nums, target)
		r.Equal(result, expectedResult)
	})

	t.Run("test two sum", func(t *testing.T) {
		r := require.New(t)

		expectedResult := []int(nil)
		nums := []int{2, 3, 1, 11, 4, 6, 7, 24, 9}
		target := 111

		result := TwoSumMapValues(nums, target)
		r.Equal(result, expectedResult)
	})
}
