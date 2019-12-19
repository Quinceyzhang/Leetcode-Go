package main

/**
* @Author: Quincy
* @Date: 2019/11/13 9:49
* @Description: 两数相加
给出两个非空的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字0之外，这两个数都不会以0开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
*/

//Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root = ListNode{}
	var cursor = root
	var carry, sum int
	for ; l1!= nil || l2.Next != nil; {
		sum = l1.Val + l2.Val + carry
		carry = sum / 10
		var node = ListNode{0, nil}
		node.Val = sum % 10
		cursor.Next = &node
		cursor = node
		if l1.Next != nil {
			l1 = l1.Next
		}
		if l2.Next != nil {
			l2 = l2.Next
		}
	}
	return root.Next
}