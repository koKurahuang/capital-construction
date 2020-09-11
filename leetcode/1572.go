package leetcode

func diagonalSum(mat [][]int) int {
	n := len(mat)
	mid := n / 2
	var ans int
	for i:= 0 ; i< n ;i ++{
		ans += mat[i][i]
		ans += mat[i][n-1-i]
	}
	if n %2 == 0 {
		return ans
	}else {
		return ans- mat[mid][mid]
	}
}
