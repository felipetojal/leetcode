package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    nodeMap := make(map[*ListNode]*ListNode)
    
    for head.Next != nil {
        _, ok := nodeMap[head]
        if ok {
            return true
        }
        nodeMap[head] = head.Next
        head = head.Next
    }
    return false
}

