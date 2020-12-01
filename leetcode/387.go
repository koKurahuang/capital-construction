package leetcode

func firstUniqChar(s string) int {
	cntMap  := make([]int,26)
	locMap := make([]int,26)
	for i, k := range s {
		cntMap[int(k)-97] ++
		locMap[int(k)-97] = i
	}
	for _, k := range s {
		if cntMap[int(k)-97] == 1 {
			return locMap[int(k)-97]
		}
	}
	return -1
}
