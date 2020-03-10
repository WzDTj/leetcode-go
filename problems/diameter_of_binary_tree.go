package problems

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max int
	traverse(&max, root)
	return max
}

func traverse(max *int, root *TreeNode) int {
	var leftDeep, rightDeep, rootDeep, radius int

	if root.Left != nil {
		leftDeep = traverse(max, root.Left)
	}
	if root.Right != nil {
		rightDeep = traverse(max, root.Right)
	}

	radius = leftDeep + rightDeep
	*max = int(math.Max(float64(*max), float64(radius)))

	rootDeep = 1 + int(math.Max(float64(leftDeep), float64(rightDeep)))

	return rootDeep
}
