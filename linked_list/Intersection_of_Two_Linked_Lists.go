package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var s []*ListNode

	for headA != nil {
		s = append(s, headA)
		headA = headA.Next
	}
	for headB != nil {
		for _, v := range s {
			if v == headB {
				return headB
			}
		}
		headB = headB.Next
	}
	return nil
}

/**
 * 链表拼接法
 */
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

/**
 * 消除链表长度差
 */
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA := getLen(headA)
	lenB := getLen(headB)

	//计算链表长度差 n，长的先移动 n 步
	if lenA > lenB { // 链表A比链表B长，A先移动
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	} else { // 链表B比链表A长，B先移动
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

//获取链表长度
func getLen(head *ListNode) int {
	var len int
	for head != nil {
		len++
		head = head.Next
	}
	return len
}
