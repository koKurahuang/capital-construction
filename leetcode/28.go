package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkMirror(root.Left, root.Right)
}

func checkMirror(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return checkMirror(l.Left,r.Right ) && checkMirror(l.Right, r.Left )
}