package leetcode

// 反转链表
func reverseListMethod001(head *ListNode) *ListNode {
	var prev *ListNode
	var tmp *ListNode
	cur := head

	for cur != nil {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

