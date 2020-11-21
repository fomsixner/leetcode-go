package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//链表排序, O(nlog(n))
//归并排序,自底向上
func merge(head1, head2 *ListNode) *ListNode {
	newhead := &ListNode{}
	temp := newhead
	ptr1, ptr2 := head1, head2
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			temp.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			temp.Next = ptr2
			ptr2 = ptr2.Next
		}
		temp = temp.Next
	}
	if ptr1 != nil {
		temp.Next = ptr1
	} else if ptr2 != nil {
		temp.Next = ptr2
	}
	return newhead.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for ptr := head; ptr != nil; ptr = ptr.Next { //求链表长度
		length++
	}
	newhead := &ListNode{Next: head}
	for sublength := 1; sublength < length; sublength *= 2 { // log(n)次循环
		pre, cur := newhead, newhead.Next
		for cur != nil {
			//断链
			head1 := cur
			for i := 1; i < sublength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil //获得第一个链表
			cur = head2
			for i := 1; i < sublength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			//归并
			pre.Next = merge(head1, head2)
			for pre.Next != nil {
				pre = pre.Next
			}
			cur = next
		}
	}
	return newhead.Next
}
