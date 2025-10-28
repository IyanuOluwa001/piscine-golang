package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

func isBST(node, min, max *TreeNode) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {
		return false
	}

	return isBST(node.Left, min, node) && isBST(node.Right, node, max)
}
