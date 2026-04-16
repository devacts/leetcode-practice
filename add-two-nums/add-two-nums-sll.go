package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, carry := 0, 0
	current := &ListNode{}
	head := current
	for l1 != nil || l2 != nil || carry > 0 {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		sum = sum % 10

		current.Val = sum
		current.Next = &ListNode{Next: nil}
		if l1 == nil && l2 == nil && carry == 0 {
			current.Next = nil
		}
		current = current.Next
	}
	return head
}

func main() {
	// l1 = [2,4,3]
	l14 := ListNode{Val: 9}
	l13 := ListNode{Val: 3, Next: &l14}
	// l13 := ListNode{Val: 3}
	l12 := ListNode{Val: 4, Next: &l13}
	l11 := ListNode{Val: 2, Next: &l12}

	// l2 = [5,6,4]
	l23 := ListNode{Val: 4}
	l22 := ListNode{Val: 6, Next: &l23}
	l21 := ListNode{Val: 5, Next: &l22}
	newNode := addTwoNumbers(&l11, &l21)
	fmt.Printf("new node is: %v", newNode)
}
