package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//时间复杂度：O(n)，其中 n 是链表的长度。需要遍历链表一次。
//空间复杂度：O(1)。
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev //
		prev = curr
		curr = next
	}
	return prev
}

//时间复杂度：O(n)，其中 n 是链表的长度。需要对链表的每个节点进行反转操作。
//空间复杂度：O(n)，其中 n 是链表的长度。空间复杂度主要取决于递归调用的栈空间，最多为 n 层。
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next) //返回时，newHead指向尾节点，递归调用栈中保存着旧链表中的每一个节点head
	head.Next.Next = head              //head为即将加入新链表的节点，head.Next为新链表的尾节点
	head.Next = nil
	return newHead
}

func printNode(head *ListNode) {
	for head != nil {
		//fmt.Print(head.value, "\t")
		fmt.Println(head.Val)
		head = head.Next
	}
}

//判断链表是否存在环
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
