package p001x

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func SliceToChain(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var head *ListNode = &ListNode{nums[0], nil}
	for i, le, tail := 1, len(nums), head; i < le; i++ {
		tail.Next = &ListNode{nums[i], nil}
		tail = tail.Next
	}
	return head
}

func ChainToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var nums []int = []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}
