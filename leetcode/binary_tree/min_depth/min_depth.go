// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// https://www.youtube.com/watch?v=hmWhJyz5kqc
package min_depth

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Depth first traversal
// O(n) time
// O(log n) space
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
