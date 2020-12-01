package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	one := binary(nums, target)
	if one == -1 {
		return []int{-1, -1}
	}
	var ret = make([]int,2)
	for {
		if one +1 == len(nums) || nums[one+1] != target{
			ret[1] = one
			break
		}else{
			one ++
		}

	}
	for {
		if one - 1 < 0 ||  nums[one-1] != target{
			ret[0] = one
			break
		}else{
			one --
		}

	}
	return ret
}

func binary(nums []int , target int) int {
	head := 0
	tail := len(nums)

	for {
		mid := (head +tail) /2
		if nums[mid] == target {
			return mid
		}
		if mid <= head || mid >= tail {
			return -1
		}
		if nums[mid] < target {
			head = mid + 1
			continue
		}
		if nums[mid] > target {
			tail = mid -1
			continue
		}
	}

}
