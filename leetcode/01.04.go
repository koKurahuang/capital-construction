package leetcode

func canPermutePalindrome(s string) bool {
	cnt := make(map[int]int, 0)
	for _,v := range s {
		if cnt[int(v)] == 0 {
			cnt[int(v)] = 1
		}else {
			cnt[int(v)] = 0
		}
	}
	var sum int
	for _,v := range cnt {
		sum += v
	}
	if sum > 1 {
		return false
	}else {
		return true
	}
}