package main

// expected boolean expression, found assignment (missing parentheses around composite literal?)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 当只有一个元素时
	if head.Next == nil {
		return true
	}

	// 有两个元素及以上
	var a []int
	temp := head

	for temp != nil {
		a = append(a, temp.Val)
		temp = temp.Next
	}

	lenth := len(a)
	for i := 0; i < lenth/2; i++ {
		if a[i] != a[lenth-i-1] {
			return false
		}
	}

	return true
}
