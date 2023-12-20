package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	head = reverse(head)
	if n == 1 {
		return reverse(head.Next)
	}
	first := head
	var last *ListNode
	for i := 1; i < n; i++ {
		last = first
		first = first.Next
	}
	last.Next = first.Next
	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	first, second := head, head.Next
	first.Next = nil
	for second != nil {
		temp := second.Next
		second.Next = first
		first = second
		second = temp
	}
	return first
}
