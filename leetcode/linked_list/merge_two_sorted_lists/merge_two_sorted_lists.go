// https://leetcode.com/problems/merge-two-sorted-lists/
// https://www.youtube.com/watch?v=bdWOmYL5d1g
package merge_two_sorted_lists

// Input: l1 = [1,2,4], l2 = [1,3,4]
// Output: [1,1,2,3,4,4]

// L				   L							L							 L					L
// 1, 2, 4	=> 1, 2, 4	=> 1, 2, 4	=> 1, 2, 4 => 1, 2, 4

// R							R						R						R							R						 R
// 1, 3, 4 	=> 1, 3, 4	=> 1, 3, 4	=> 1, 3, 4 => 1, 3, 4 => 1, 3, 4

// 															 T
// Starting Random Dummy Node = -1

//     T		 		   T 				   	  T		   			      T							     	 T										   T
// -1, 1 => -1, 1, 1 => -1, 1, 1, 2 => -1, 1, 1, 2, 3 => -1, 1, 1, 2, 3, 4 => -1, 1, 1, 2, 3, 4, 4

// Return first node.next

// Node in singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// O(m + n)
// O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
