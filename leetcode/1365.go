package leetcode

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	var a []int
	for i,_ := range nums {
		a = append(a, nums[i])
	}
	sort.Ints(a)

	var record = make(map[int]int)
	for i,_ := range a {
		if _, ok := record[a[i]]; !ok {
			record[a[i]] = i
		}
	}
	var res []int
	for i, _ := range nums {
		res = append(res, record[nums[i]])
	}
	return res
}
