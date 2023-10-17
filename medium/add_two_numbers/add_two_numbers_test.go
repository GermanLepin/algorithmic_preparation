package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AddTwoNumbers(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		r := require.New(t)

		// l1
		var (
			l13 = &ListNode{
				Val:  6,
				Next: nil,
			}

			l12 = &ListNode{
				Val:  9,
				Next: l13,
			}

			l1 = &ListNode{
				Val:  7,
				Next: l12,
			}
		)

		// l2
		var (
			l23 = &ListNode{
				Val:  9,
				Next: nil,
			}

			l22 = &ListNode{
				Val:  8,
				Next: l23,
			}

			l2 = &ListNode{
				Val:  7,
				Next: l22,
			}
		)

		var (
			expectedResult14 = &ListNode{
				Val:  1,
				Next: nil,
			}

			expectedResult13 = &ListNode{
				Val:  6,
				Next: expectedResult14,
			}

			expectedResult12 = &ListNode{
				Val:  8,
				Next: expectedResult13,
			}

			expectedResult = &ListNode{
				Val:  4,
				Next: expectedResult12,
			}
		)

		result := AddTwoNumbers(l1, l2)
		r.Equal(result, expectedResult)
	})
}
