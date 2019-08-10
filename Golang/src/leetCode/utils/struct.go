package utils

type ListNode struct {
    Val  int
    Next *ListNode
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
