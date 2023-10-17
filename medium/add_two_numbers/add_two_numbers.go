package add_two_numbers

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Leetcode https://leetcode.com/problems/add-two-numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head            = new(ListNode)
		currentListNode = head
		carry           = 0
	)

	for l1 != nil || l2 != nil || carry != 0 {
		var (
			value1 = 0
			value2 = 0
		)

		if l1 != nil {
			value1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			value2 = l2.Val
			l2 = l2.Next
		}

		//  987
		//  234
		// 1221

		// 3+5 = 8 and carry = 0, if we have 7+8 = 15, carry = 1 (max is 18(9+9)) carry can be 1 or 0
		// 2+3 = 5 + 1 = 6
		num := value1 + value2 + carry
		carry = num / 10 // 9/10 = 0, 11/10 = 1

		currentListNode.Next = &ListNode{Val: num % 10, Next: nil} // % - remainder operator 18 % 10 = 8, 9 % 10 = 9
		currentListNode = currentListNode.Next
	}

	return head.Next
}
