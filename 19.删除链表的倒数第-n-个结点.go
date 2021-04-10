package leetcode_go

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil { // 删除的是唯一节点
		return nil
	}
	left, right := head, head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	if right == nil { // 删除的是第一个节点
		head = head.Next
	} else {
		for right.Next != nil {
			left = left.Next
			right = right.Next
		}
		left.Next = left.Next.Next
	}
	return head
}

// @lc code=end
