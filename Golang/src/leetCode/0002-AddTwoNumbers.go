/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetCode

import "fmt"

func Main0002() {
	l1 := &listNode{0, &listNode{1, nil}}
	l2 := &listNode{3, &listNode{4, nil}}
	fmt.Println("Input: ")
	myprint(l1)
	fmt.Println()
	myprint(l2)
	fmt.Println()
	result := addTwoNumbers(l1, l2)
	fmt.Println("Output: ")
	myprint(result)
}

type listNode struct {
	Val  int
	Next *listNode
}

func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	numUp := 0
	newHead := &listNode{0, nil}
	newEnd := newHead
	for (l1 != nil) || (l2 != nil) || (numUp != 0) {
		if l1 != nil {
			numUp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			numUp += l2.Val
			l2 = l2.Next
		}
		if numUp > 9 {
			newEnd.Next = &listNode{numUp - 10, nil}
			numUp = 1
		} else {
			newEnd.Next = &listNode{numUp, nil}
			numUp = 0
		}
		newEnd = newEnd.Next
	}
	return newHead.Next
}

func myprint(l *listNode) {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
}
