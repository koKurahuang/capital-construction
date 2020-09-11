package leetcode

import (
	"fmt"

)

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	for i:=1;i<=k;i++{
		if k-i == len(num){
			return "0"
		}
		if num[0]> num[1] {
			num = num[1:]
			fmt.Println("a",i,"  ",num)
			continue
		}
		for j:=1;j<len(num);j++{

			if num[j-1]>num[j] {
				if j == 1 {
					num = num[1:]
				}else{
					num = num[:j-1]+num[j:]
				}
				fmt.Println("b",i,num)
				break
			}
			if j == len(num) - 1 {
				num = num[:len(num)-1]
				break
			}
		}
		fmt.Println("c",i,num)
	}

	return clear(num)
}

func clear(s string) string {
	ss := s
	for i:=0;i<len(s);i++{
		if string(s[i]) == "0" {
			ss = s[i+1:]
		}else{
			break
		}
	}
	if ss == "" {
		ss = "0"
	}
	return ss
}

