package leetcode_mayulu

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = nil
	advance := 0

	cl1 := l1
	cl2 := l2
	var previous *ListNode = nil
	v1 := 0
	v2 := 0
	sum := 0
	for cl1 != nil || cl2 != nil || advance != 0 {
		if cl1 != nil {
			v1 = cl1.Val
		} else {
			v1 = 0
		}
		if cl2 != nil {
			v2 = cl2.Val
		} else {
			v2 = 0
		}
		sum = v1 + v2 + advance
		current := new(ListNode)
		current.Val = sum % 10
		current.Next = nil
		if result == nil {
			previous = current
			result = current
		} else {
			previous.Next = current
			previous = current
		}
		advance = sum / 10
		if cl1 != nil {
			cl1 = cl1.Next
		}
		if cl2 != nil {
			cl2 = cl2.Next
		}
	}
	return result
}
