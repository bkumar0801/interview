/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

/**
 * Definition for singly-linked list.
 */
 type ListNode struct {
      Val int
      Next *ListNode
 }
 

func NewListNode(v int) *ListNode {
    return &ListNode{
        Val:v,
        Next: nil,
    }
}

func Empty(l *ListNode) bool {
    return l == nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if Empty(l1) && Empty(l2) {
        return nil
    }
    if Empty(l1) {
        return l2
    } else if Empty(l2) {
        return l1
    }
    
    h1 := l1
    h2 := l2
    carry := 0
    result := NewListNode(-1)
    iter := result
    for h1 != nil || h2 != nil {
        v1 := 0
        v2 := 0
        if !Empty(h1) {
            v1 = h1.Val
        }
        if !Empty(h2) {
            v2 = h2.Val
        }
        sum := v1 + v2 + carry
        iter.Next = NewListNode(sum % 10)
        carry = sum / 10
        if !Empty(h1) {
            h1 = h1.Next
        }
        if !Empty(h2) {
            h2 = h2.Next
        }
        iter = iter.Next
    }
    if carry != 0 {
        iter.Next = NewListNode(carry)
    }
    return result.Next
}
