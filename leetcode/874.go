package leetcode

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {

	max := 0
	var nowX, nowY = 0, 0
	var basicX, basicY = 0, 1
	var recordMap = make(map[[2]int]bool)
	for i, _ := range obstacles {
		recordMap[[2]int{obstacles[i][0],obstacles[i][1]}] =true
	}
	var left = func() {
		if basicX == 0 {
			if basicY == 1 {
				basicX, basicY = -1, 0
			} else {
				basicX, basicY = 1, 0
			}
		} else if basicX == 1 {
			basicX, basicY = 0, 1
		} else {
			basicX, basicY = 0, -1
		}

	}

	var right = func() {
		if basicX == 0 {
			if basicY == 1 {
				basicX, basicY = 1, 0
			} else {
				basicX, basicY = -1, 0
			}
		} else if basicX == 1 {
			basicX, basicY = 0, -1
		} else {
			basicX, basicY = 0, 1
		}
	}

	for _, v := range commands {
		if v > 0 {
			for i := 1; i <= v; i++ {
				targetX, targetY := nowX+1*basicX, nowY+1*basicY
				if ifNoPass(targetX, targetY, recordMap) {
					break
				} else {
					nowX, nowY = targetX, targetY
					if nowX*nowX+nowY*nowY > max {
						max = nowX*nowX + nowY*nowY
					}
				}
			}
		} else if v == -2 {
			left()
		} else {
			right()
		}

		fmt.Println(nowX, nowY)
	}

	return max
}

func ifNoPass(x, y int, obs map[[2]int]bool) bool {
	if _, ok := obs[[2]int{x,y}]; ok {
		return true
	}
	return false
}
