/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	if carry := helper(head); carry > 0 {
		head = &ListNode{carry, head}
	}
	return head
}

func helper(head *ListNode) int {
	var carry int
	if head.Next == nil {
		head.Val, carry = double(head.Val, 0)
		return carry
	}
	head.Val, carry = double(head.Val, helper(head.Next))
	return carry
}

func double(val, adder int) (int, int) {
	val *= 2
	val += adder
	return val % 10, val / 10
}
