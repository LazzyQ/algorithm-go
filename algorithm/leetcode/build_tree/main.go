package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeWithPreorderAndInorder(preorder, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	i := 0
	root := preorder[0]
	for ; i < len(inorder); i++ {
		if root == inorder[i] {
			break
		}
	}

	h := &TreeNode{
		Val: root,
	}

	if i == 0 {
		h.Right = buildTreeWithPreorderAndInorder(preorder[1:], inorder[i+1:])
	} else if i == len(inorder)-1 {
		h.Left = buildTreeWithPreorderAndInorder(preorder[1:], inorder[:i])
	} else {
		h.Left = buildTreeWithPreorderAndInorder(preorder[1:i+1], inorder[:i])
		h.Right = buildTreeWithPreorderAndInorder(preorder[i+1:], inorder[i+1:])
	}
	return h
}

func buildTreeWithInorderAndPostorder(inorder, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	root := postorder[len(postorder)-1]
	i := 0
	for ; i < len(inorder); i++ {
		if root == inorder[i] {
			break
		}
	}

	h := &TreeNode{Val: root}
	if i == len(postorder)-1 {
		h.Left = buildTreeWithInorderAndPostorder(inorder[:i], postorder[:i])
	} else if i == 0 {
		h.Right = buildTreeWithInorderAndPostorder(inorder[1:], postorder[:len(postorder)-1])
	} else {
		h.Left = buildTreeWithInorderAndPostorder(inorder[:i], postorder[:i])
		h.Right = buildTreeWithInorderAndPostorder(inorder[i+1:], postorder[i:len(postorder)-1])
	}
	return h
}
