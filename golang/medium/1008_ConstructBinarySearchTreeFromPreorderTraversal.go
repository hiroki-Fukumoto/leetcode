package leetcode

// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/
// Runtime: 2 ms
// Memory Usage: 2.3 MB

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := 1
	for i < len(preorder) && preorder[i] < preorder[0] {
		i++
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:i]),
		Right: bstFromPreorder(preorder[i:]),
	}
}
