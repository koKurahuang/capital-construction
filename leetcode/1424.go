package leetcode
// 你妈的  54测试用例超时了
func findDiagonalOrder(nums [][]int) []int {
	chang := 0
	for i, _ := range nums {
		if len(nums[i]) > chang {
			chang = len(nums[i])
		}
	}
	gao := len(nums)
	max := 0
	if chang > gao {
		max = chang
	}else{
		max = gao
	}

	var output []int
	var i, j int
	for k:=0;k < 2*max -1;k++ {
		if k<max {
			i = k
			j = 0
		}else{
			i = max -1
			j = k-max + 1
		}


		for i >= 0 {
			if i >= gao {
				i --
				j++
				continue
			}
			if j < len(nums[i]) {
				output = append(output, nums[i][j])
			}

			j++
			i--
		}

	}
	return output
}
