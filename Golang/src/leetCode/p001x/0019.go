package p001x

import . "leetCode/utils"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 题目规定“给定的 n 保证是有效的。”所以不对n进行检查了
	// 这里直接进阶：你能尝试使用一趟扫描实现吗？
	var length int = n + 1
	var tempNodes []*ListNode = make([]*ListNode, length)
	var countNode int = 0
	var tail *ListNode = head
	for tail != nil {
		tempNodes[countNode%length] = tail
		tail = tail.Next
		countNode++
	}
	if countNode == n { // 最后一个节点的情况
		return head.Next
	}
	if n == 1 { // 第一个节点的情况
		tempNodes[countNode%length].Next = nil
	} else { // 中间的情况
		tempNodes[countNode%length].Next = tempNodes[(countNode+2)%length]
	}
	return head
}

