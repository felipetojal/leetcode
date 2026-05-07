package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := &ListNode{}
	aux := list
	for list1 != nil && list2 != nil {
		fmt.Println("--------------")	
		if list1.Val <= list2.Val {
			fmt.Printf("list1.Val: %d <= list2.Val: %d\n", list1.Val, list2.Val)
			aux.Next = list1
			fmt.Printf("aux.Val: %d\n", list1.Val)
			list1 = list1.Next
		} else  {
			fmt.Printf("list1.Val: %d > list2.Val: %d\n", list.Val, list2.Val)
			aux.Next = list2
			fmt.Printf("aux.Val: %d\n", list2.Val)
			list2 = list2.Next
		}

		aux = aux.Next
	}

	if list1 == nil {
		aux.Next = list2
	} else if list2 == nil {
		aux.Next = list1
	}

	return list.Next
}