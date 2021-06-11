package leetcode_go

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	merge := func(A, B *ListNode) *ListNode {
		head := &ListNode{}
		ptr := head
		ptr1, ptr2 := A, B
		for ptr1 != nil && ptr2 != nil {
			if ptr1.Val < ptr2.Val {
				ptr.Next = ptr1
				ptr1 = ptr1.Next
			} else {
				ptr.Next = ptr2
				ptr2 = ptr2.Next
			}
			ptr = ptr.Next
		}
		if ptr1 != nil {
			ptr.Next = ptr1
		} else {
			ptr.Next = ptr2
		}
		return head.Next
	}
	var mergeList func(start, end int) *ListNode
	mergeList = func(start, end int) *ListNode {
		if start == end {
			return lists[start]
		}
		if end-start == 1 {
			return merge(lists[start], lists[end])
		}
		A := mergeList(start, (start+end)/2)
		B := mergeList((start+end)/2+1, end)
		return merge(A, B)
	}
	return mergeList(0, len(lists)-1)
}

// @lc code=end
