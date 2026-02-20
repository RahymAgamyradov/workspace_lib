package BinaryTrees

import "fmt"

// TreeNode - a struct that takes nodes of a binary tree.
//
//	It has a value and left and right children.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ScanTree - a function that you can scan Binary trees
// You just first take a TreeNode [a] and write [a.ScanTree()]
func (root *TreeNode) ScanTree() {
	var n int
	fmt.Scan(&n)
	a := make([]TreeNode, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i].Val)
	}
	for i := 1; i <= n; i++ {
		if i*2 <= n {
			if a[i*2].Val != 0 {
				a[i].Left = &a[i*2]
			}
		}
		if i*2+1 <= n {
			if a[i*2+1].Val != 0 {
				a[i].Right = &a[i*2+1]
			}
		}
	}
	*root = a[1]
}
