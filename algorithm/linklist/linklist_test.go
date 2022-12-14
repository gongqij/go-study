package main

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	fmt.Println("before reverse:")
	printNode(node1)
	head := reverseList(node1) //复制了一份指针(node1)，执行结束node1还是指向同一个ListNode
	fmt.Println("after reverse:")
	printNode(head)
}
func TestHasCycle(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(hasCycle(node1))
}
