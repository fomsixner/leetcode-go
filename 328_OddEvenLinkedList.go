package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, odd := head, head.Next.Next
	temp := head.Next
	for odd != nil {
		temp.Next = odd.Next
		temp = odd.Next
		odd.Next = cur.Next
		cur.Next = odd
		cur = cur.Next
		if temp != nil {
			odd = temp.Next
		} else {
			break
		}
	}
	return head
}
