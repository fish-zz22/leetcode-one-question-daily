package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	k := 3
	head := new(ListNode)
	dumpy := new(ListNode)
	dumpy.Next = head
	for i, v := range nums {

		if i == len(nums)-1 {
			break
		}

		head.Val = v
		head.Next = new(ListNode)
		head = head.Next
	}
	splitListToParts(dumpy.Next, k)
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	listNodes := make([]*ListNode, k)
	var l, index int
	temp, cur := head, head
	tempK := k

	for temp != nil {
		l++
		temp = temp.Next
	}

	for cur != nil && index < k {
		// 获取链表需要向后移的步数
		step := int(math.Ceil(float64(l) / float64(tempK)))

		// 第k个链表进行赋值
		listNodes[index] = cur

		// 链表需要向后移step-1
		for i := 1; i < step; i++ {
			cur = cur.Next
		}

		// l 更新为总长度-step
		l = l - step
		tempK--
		index++
		cur, cur.Next = cur.Next, nil
	}
	return listNodes
}
