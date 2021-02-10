// https://leetcode.com/problems/maximum-depth-of-binary-tree/
package max_depth

// binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n) time because we need to traverse every node on the tree
// O(log n) space because the call stack can only get as big as the depth
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
