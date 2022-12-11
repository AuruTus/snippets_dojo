package lc_snippets

import (
	"context"
	snippets "snippets_dojo/src"
)

type LC98Tstr struct{}

var _ snippets.Tstr = (*LC98Tstr)(nil)

func (t *LC98Tstr) Test(ctx context.Context) error {
	isValidBST(nil)
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func validateSubtree(root *TreeNode) (isValid bool, maxValNode, minValNode *TreeNode) {
	if root == nil {
		return true, nil, nil
	}
	lValid, lMaxValNode, lMinValNode := validateSubtree(root.Left)
	rValid, rMaxValNode, rMinValNode := validateSubtree(root.Right)

	maxValNode = rMaxValNode
	if rMaxValNode == nil {
		maxValNode = root
	}
	minValNode = lMinValNode
	if lMinValNode == nil {
		minValNode = root
	}
	/**
	 * ? we can just validate the order between root node and its left and right sub trees (interval order),
	 *  which means there's **no need** to compare the root's value with its direct children's
	 *  value. I guess this can be proved recursively?
	 *  */
	isValid = lValid && ((lMaxValNode != nil && lMaxValNode.Val <= root.Val) || lMaxValNode == nil) &&
		// ((root.Left != nil && root.Left.Val <= root.Val) || root.Left == nil) &&
		rValid && ((rMinValNode != nil && rMinValNode.Val >= root.Val) || rMinValNode == nil)
	// ((root.Right != nil && root.Right.Val >= root.Val) || root.Right == nil)
	return
}

func isValidBST(root *TreeNode) bool {
	isValid, _, _ := validateSubtree(root)
	return isValid
}
