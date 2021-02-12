// https://leetcode.com/problems/balanced-binary-tree/
// https://www.youtube.com/watch?v=LU4fGD-fgJQ

package balanced_binary_tree

// binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
- Continuously ask the left node for its height & then ask the right
- if abs(depth of left node - depth of right node) <= 1 then current node is balanced
*/
// O(N) time where n is the # of nodes in the tree
// O(H) space where h is the height of the tree
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return abs(leftHeight-rightHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
