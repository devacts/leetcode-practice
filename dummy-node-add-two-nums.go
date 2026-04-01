package main

import "fmt"

// Definition for singly-linked list.
type DListNode struct {
	Val  int
	Next *DListNode
}

func addTwoNumbersd(l1 *DListNode, l2 *DListNode) *DListNode {
	dummy := &DListNode{} // The starting point of our result list
	current := dummy      // A pointer to move as we add nodes
	carry := 0

	// Continue as long as there is data in either list or a carry remains
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next // Move to next node in l1
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next // Move to next node in l2
		}

		// Update carry for the next position
		carry = sum / 10

		// Create the new node with the single digit (remainder)
		current.Next = &DListNode{Val: sum % 10}

		// Move our result pointer forward
		current = current.Next
	}

	return dummy.Next // Return the first real node (skipping our dummy)
}

func main() {
	// l1 = [2,4,3]
	l14 := DListNode{Val: 9}
	l13 := DListNode{Val: 3, Next: &l14}
	// l13 := DListNode{Val: 3}
	l12 := DListNode{Val: 4, Next: &l13}
	l11 := DListNode{Val: 2, Next: &l12}

	// l2 = [5,6,4]
	l23 := DListNode{Val: 4}
	l22 := DListNode{Val: 6, Next: &l23}
	l21 := DListNode{Val: 5, Next: &l22}
	newNode := addTwoNumbersd(&l11, &l21)
	fmt.Printf("new node is: %v", newNode)
}
