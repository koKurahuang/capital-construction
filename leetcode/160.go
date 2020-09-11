package leetcode

import "strconv"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	astr, acnt := reversString(headA)
	bstr, _ := reversString(headB)
	var record int
	for i:=1; i<=acnt; i++ {
		if string(astr[i]) != string(bstr[i]) {
			record = i
			break
		}
	}
	alist := headA
	for record != 0 {
		record --
		alist = headA.Next
	}
	return alist
}

func reversString(head *ListNode) (string, int) {
	var s string
	var cnt int
	for head.Next != nil {
		s = strconv.Itoa(head.Val) + s
		head = head.Next
		cnt ++
	}
	return s, cnt
}