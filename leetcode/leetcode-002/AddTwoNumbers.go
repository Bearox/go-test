package main

func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	valTmp := (l1.Val + l2.Val) % 10
	nextPlus := (l1.Val + l2.Val) / 10
	retHead := &ListNode{
		Val: valTmp,
	}
	previous := retHead
	for l1.Next != nil || l2.Next != nil || nextPlus == 1 {
		v1 := 0
		v2 := 0
		if l1.Next != nil {
			l1 = l1.Next
			v1 = l1.Val
		}
		if l2.Next != nil {
			l2 = l2.Next
			v2 = l2.Val
		}
		valTmp = (v1 + v2 + nextPlus) % 10
		nextPlus = (v1 + v2 + nextPlus) / 10
		currentNode := &ListNode{
			Val: valTmp,
		}
		previous.Next = currentNode
		previous = currentNode
	}
	return retHead
}