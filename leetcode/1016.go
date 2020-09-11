package leetcode

/*
执行用时：
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗：
2.1 MB
, 在所有 Go 提交中击败了
25.00%
的用户
 */

import (
	"fmt"
	"strconv"
	"strings"
)
func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		if !strings.Contains(S, bin(i)) {
			return false
		}
	}
	return true
}
func bin(i int) string {
	var s string
	left := i
	for {
		if left > 0 {
			s = strconv.Itoa(left % 2) + s
			left = left/2
		} else {
			//s += strconv.Itoa(left % 2)
			break
		}
	}
	fmt.Printf("%d ->   %s\n", i, s)
	return s
}
