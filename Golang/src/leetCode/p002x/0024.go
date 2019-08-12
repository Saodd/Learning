package p002x

import (
    . "leetCode/utils"
)

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var fakeHead *ListNode = &ListNode{Next: head}
    var fNode *ListNode = fakeHead
    var lNode, rNode *ListNode
    for {
        if head == nil {
            break
        }
        lNode = head

        head = head.Next
        if head == nil {
            break
        }
        rNode = head

        head = head.Next
        fNode.Next, rNode.Next, lNode.Next = rNode, lNode, head
        fNode = lNode

    }
    return fakeHead.Next
}
