/*
狗屎题目，题目误解老子，看了半天示例3
 */
package leetcode

func longestSubsequence(arr []int, difference int) int {
	max := 0
	var dp = make(map[int]int)
	for i:=0;i<len(arr);i++ {
		dp[arr[i]]= dp[arr[i]-difference]+1
		if dp[arr[i]] > max {
			max = dp[arr[i]]
		}
	}
	return max
}
