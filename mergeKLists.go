package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var ans *ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	ans = group(lists)
	return ans
}

func group(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 1 {
		return lists[0]
	}
	return mergetwoLists(group(lists[:n/2]), group(lists[n/2:]))
}

func mergetwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
