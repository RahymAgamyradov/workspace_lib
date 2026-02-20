package LinkedLists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (root *ListNode) ScanList() {
	var size int
	fmt.Scan(&size)
	ans := make([]ListNode, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&ans[i].Val)
	}
	for i := 0; i < size-1; i++ {
		ans[i].Next = &ans[i+1]
	}
	*root = ans[0]
}
func (root ListNode) PrintList() {
	current := &root
	for current != nil {
		fmt.Printf("%v ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
