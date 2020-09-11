package leetcode

import (
	"fmt"
)

func patternMatching(pattern string, value string) bool {

	var cnta, cntb int
	for i, _ := range pattern {
		if pattern[i] == 'a' {
			cnta++
		}
	}
	cntb = len(pattern) - cnta

	if value == "" {
		return cnta==0 || cntb==0
	}
	if pattern == "" {
		return value == ""
	}

	res := figureXY(cnta, cntb, len(value))
	fmt.Println("cnt", cnta, cntb)
	fmt.Println("figure", res)
	if len(res) == 0 {
		fmt.Println("1 false")
		return false
	}
	for i, _ := range res {
		ans := res[i]
		x, y := ans[0], ans[1]
		if x == 0 && y == 0 {
			if len(value) == 0 {
				fmt.Println("1 true")
				return true
			} else {
				fmt.Println("2 false")
				return false
			}
		}
		if x == 0 {
			stry := value[:y]
			wantedy := ""
			for cnti := 1; cnti <= cntb; cnti++ {
				wantedy += stry
			}
			if wantedy == value {
				fmt.Println("2 true")
				return true
			} else {
				continue
			}
		}
		if y == 0 {
			strx := value[:x]
			wantedx := ""
			for cnti := 1; cnti <= cnta; cnti++ {
				wantedx += strx
			}
			if wantedx == value {
				fmt.Println("3 true")
				return true
			} else {
				continue
			}
		}
		wantedstr := ""

		var stra, strb string
		if pattern[0] == 'a' {
			stra = value[0:x]
			cnt := 0
			for k := 0; k < len(pattern); k++ {
				if pattern[k] == 'a' {
					wantedstr += stra
					cnt += x
				} else {
					if strb == "" {
						strb = value[cnt : cnt+y]
						if stra == strb {
							continue
						}
					}
					wantedstr += strb
					cnt += y
				}
			}
		} else {
			strb = value[0:y]
			cnt := 0
			for k := 0; k < len(pattern); k++ {
				if pattern[k] == 'b' {
					wantedstr += strb
					cnt += y
				} else {
					if stra == "" {
						stra = value[cnt : cnt+x]
						if stra == strb {
							continue
						}
					}
					wantedstr += stra
					cnt += x
				}
			}
		}
		if wantedstr == value {
			fmt.Println("4 true")
			return true
		}
	}
	fmt.Println("5 false")
	return false

}

func figureXY(a, b int, result int) ([][2] int) {
	// ax + by = result
	var res [][2]int
	if a == 0 {
		if b == 0 {
			return nil
		}else {
			if result % b == 0 {
				res =append(res, [2]int{0, result / b})
				return res
			}else {
				return res
			}
		}
	}
	if b == 0 {
		if result % a == 0 {
			res =append(res, [2]int{result / a, 0})
			return res
		}else {
			return res
		}
	}
	for x := 0; x*a <= result; x++ {
		left := result - x*a
		if b == 0 {
			continue
		}
		if left%b == 0 {
			y := left / b
			res = append(res, [2]int{x, y})
		}
	}
	return res
}
