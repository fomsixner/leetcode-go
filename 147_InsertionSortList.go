package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var res = new(ListNode)
	cur := head.Next
	res.Next, head.Next = head, nil
	for cur != nil {
		ptr := res
		for ; ptr.Next != nil && ptr.Next.Val < cur.Val; ptr = ptr.Next {
		}
		temp := cur.Next
		cur.Next = ptr.Next
		ptr.Next = cur
		cur = temp
	}
	return res.Next
}
