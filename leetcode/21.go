package leetcode

/*
偷鸡摸狗失败效率贼差
执行用时：
4 ms
, 在所有 Go 提交中击败了
61.48%
的用户
内存消耗：
2.9 MB
, 在所有 Go 提交中击败了
9.09%
的用户
 */

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := getList(l1)
	list2 := getList(l2)
	var list3 []int
	list3 = append(list3, list1...)
	list3 = append(list3, list2...)
	sort.Ints(list3)
	if len(list3) == 0{
		return nil
	}
	var start = &ListNode{list3[0], nil}
	var pre = start
	for k:=1; k< len(list3);k++  {
		var now ListNode
		now.Val = list3[k]
		pre.Next = &now
		pre = &now
	}
	return start
}

func getList(l *ListNode) []int {
	var a []int
	for l != nil {
		a = append(a, l.Val)
		l = l.Next
	}
	return a
}
