/*
执行结果：
通过
显示详情
执行用时：
456 ms
, 在所有 Go 提交中击败了
5.71%
的用户
内存消耗：
6.5 MB
, 在所有 Go 提交中击败了
5.00%
的用户
 */
package leetcode

import (
	"fmt"
	"sort"
)

func largestPerimeter(A []int) int {
	max := 0
	leng := len(A)
	sort.Ints(A)
	fmt.Println(A)
	if len(A) < 3 {
		return 0
	} else if len(A) == 3 {
		if check(A[2], A[1], A[0]) {
			return sum(A[0], A[1], A[2])
		} else {
			return 0
		}
	} else {
		for i := leng-1; i >1; i-- {
			if check(A[i], A[i-1], A[i-2]) {
				return sum(A[i], A[i-1], A[i-2])
			}
		}
	}
	return max
}

func check(a, b, c int) bool {
	if a < b+c {
		return true
	} else {
		return false
	}
}
func sum(a, b, c int) int {
	fmt.Println(a, b, c)
	return a + b + c
}
