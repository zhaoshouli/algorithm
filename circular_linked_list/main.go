package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return true
}

func createRingNodeList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	tail := head
	for i := 1; i < len(nums); i++ {
		tail.Next = &ListNode{Val: nums[i]}
		tail = tail.Next
	}
	if pos >= 0 {
		p := head
		for pos > 0 {
			p = p.Next
			pos--
		}
		tail.Next = p
	}
	return head
}

func main() {
	nums, pos := []int{3, 2, 0, -4}, 4
	head := createRingNodeList(nums, pos)
	fmt.Println(hasCycle(head))
	//nums, pos = []int{1, 2}, 0
	//head = createRingNodeList(nums, pos)
	//fmt.Println(hasCycle(head))
	//nums, pos = []int{1}, -1
	//head = createRingNodeList(nums, pos)
	//fmt.Println(hasCycle(head))
}

