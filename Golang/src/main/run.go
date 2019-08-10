package main

import (
    "fmt"
    "leetCode/utils"
)

func main() {

    fmt.Printf("%p\n",mytry())
}

func mytry() *utils.ListNode {
    head := &utils.ListNode{1, nil}
    fmt.Printf("%p\n",head)
    defer func() {head=head}()
    head = &utils.ListNode{2, nil}
    fmt.Printf("%p\n",head)
    return head
}