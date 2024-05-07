/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    if carry := helper(head); carry > 0{
        head = &ListNode{carry, head}
    }
    return head
}

func helper(head *ListNode) int {
    if head.Next == nil {
        head.Val *= 2
        newVal := head.Val % 10
        carry := head.Val / 10
        head.Val = newVal
        return carry
    }
    head.Val *= 2 
    head.Val += helper(head.Next)
    newval := head.Val % 10
    carry := head.Val / 10
    head.Val = newval
    return  carry
}
