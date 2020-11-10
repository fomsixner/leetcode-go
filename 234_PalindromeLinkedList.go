package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//判断是否是回文链表
//思路一:将数据存到列表中，用双指针比较
//思路二:使用栈
//思路三:使用递归

func isPalindrome(head *ListNode) bool {
	var list []int
	p := head
	for {
		if p == nil {
			break
		}
		list = append(list, p.Val)
		p = p.Next
	}
	var n int = len(list)
	for i := 0; i < n/2; i++ {
		if list[i] != list[n-i-1] {
			return false
		}
	}
	return true
}

//递归
func isPalindrome2(head *ListNode) bool {
	fountpointer := head
	var check func(*ListNode) bool
	check = func(curnode *ListNode) bool {
		if curnode != nil {
			if !check(curnode.Next) {
				return false
			}
			if curnode.Val != fountpointer.Val {
				return false
			}
			fountpointer = fountpointer.Next
		}
		return true
	}
	return check(head)
}
