package leetcode

import "fmt"

var cnt = 0
func findTheLongestSubstring(s string) int {
	ok := checkAEIOU(s)
	if ok {
		fmt.Println("ok       ",s)
		fmt.Println(aeiou)
		return len(s)
	}else {
		cnt ++
		fmt.Println(s)
		if cnt %2 == 0 {
			return findTheLongestSubstring(s[1:])
		}else {
			leng := len(s)
			return findTheLongestSubstring(s[:leng-1])
		}
	}

}

func checkAEIOU(s string) bool {
	aeiou = make([]int, 5)
	for _, v := range s {
		add(v)
	}
	for _, v := range aeiou {
		if v != 0 {
			return false
		}
	}
	return true
}

var aeiou = make([]int, 5)

func add(c int32) {
	switch c {
	case 'a':
		if aeiou[0] == 0 {
			aeiou[0] = 1
		} else {
			aeiou[0] = 0
		}
	case 'e':
		if aeiou[1] == 0 {
			aeiou[1] = 1
		} else {
			aeiou[1] = 0
		}
	case 'i':
		if aeiou[2] == 0 {
			aeiou[2] = 1
		} else {
			aeiou[2] = 0
		}
	case 'o':
		if aeiou[3] == 0 {
			aeiou[3] = 1
		} else {
			aeiou[3] = 0
		}
	case 'u':
		if aeiou[4] == 0 {
			aeiou[4] = 1
		} else {
			aeiou[4] = 0
		}
	}
}
