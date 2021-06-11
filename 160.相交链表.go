package leetcode_go

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 两遍扫描法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	n, m := 0, 0
	for ptr := headA; ptr != nil; ptr = ptr.Next {
		n++
	}
	for ptr := headB; ptr != nil; ptr = ptr.Next {
		m++
	}
	pa, pb := headA, headB
	if n > m {
		for i := m; i < n; i++ {
			pa = pa.Next
		}
	} else if n < m {
		for i := n; i < m; i++ {
			pb = pb.Next
		}
	}
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa, pb = pa.Next, pb.Next
	}
	return nil
}

// @lc code=end
