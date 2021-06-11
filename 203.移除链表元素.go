package leetcode_go

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	newHead := head
	for newHead != nil && newHead.Val == val {
		newHead = newHead.Next
	}
	if newHead == nil {
		return nil
	}
	ptr1, ptr2 := newHead, newHead.Next
	for ptr2 != nil {
		if ptr2.Val == val {
			ptr1.Next = ptr2.Next
			ptr2 = ptr2.Next
		} else {
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		}
	}
	return newHead
}

// @lc code=end
