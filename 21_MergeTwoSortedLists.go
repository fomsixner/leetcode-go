package leetcode_go

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newlist *ListNode = new(ListNode)
	p := l1
	q := l2
	t := newlist
	for {
		if p == nil || q == nil{
			break
		}
		temp := new(ListNode)
		if p.Val < q.Val {
			temp.Val = p.Val
			p = p.Next
		} else {
			temp.Val = q.Val
			q = q.Next
		}
		t.Next = temp
		t = t.Next
	}
	if q != nil{
		p = q
	}
	for{
		if p == nil{
			break
		}
		temp := new(ListNode)
		temp.Val = p.Val
		t.Next = temp
		p = p.Next
		t = t.Next
	}
	return newlist.Next
}

//func main()  {
//	var l1 *ListNode = new(ListNode)
//	l1.Val = -9
//	l1.Next = new(ListNode)
//	l1.Next.Val = 3
//	//l1.Next.Next = new(ListNode)
//	//l1.Next.Next.Val = 4
//	var l2 *ListNode = new(ListNode)
//	l2.Val = 5
//	l2.Next = new(ListNode)
//	l2.Next.Val = 7
//	//l2.Next.Next = new(ListNode)
//	//l2.Next.Next.Val = 4
//
//	p := l1
//	//for{
//	//	if p == nil{
//	//		break
//	//	}
//	//	fmt.Println(p.Val)
//	//	p = p.Next
//	//}
//	l3 := mergeTwoLists(l1, l2)
//	p = l3
//	for{
//		if p == nil{
//			break
//		}
//		fmt.Println(p.Val)
//		p = p.Next
//	}
//}