package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type HtNode struct {
	head *ListNode
	tail *ListNode
}

// k个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	length := getLenOfList(head)
	if length < k || k == 1 || length == 1 {
		return head
	}
	cNode := head.Next
	headList := []*HtNode{}
	pNode := head
	flag := 0
	for length >= k {
		htNode := &HtNode{tail: pNode}
		for i := 1; i < k; i++ {
			curNode := cNode.Next
			cNode.Next = pNode
			pNode = cNode
			cNode = curNode
		}
		if flag == 0 {
			head = pNode
		}
		htNode.head = pNode
		headList = append(headList, htNode)
		pNode = cNode
		if cNode != nil {
			cNode = cNode.Next
		}
		flag++
		length -= k
	}
	if length != 0 {
		htNode := &HtNode{head: pNode}
		headList = append(headList, htNode)
	} else {
		headList[len(headList)-1].tail.Next = nil
	}
	for idx, _ := range headList {
		if idx < len(headList)-1 {
			headList[idx].tail.Next = headList[idx+1].head
		}
	}
	return head
}

func getLenOfList(head *ListNode) int {
	cNode := head
	length := 0
	for cNode != nil {
		cNode = cNode.Next
		length++
	}
	return length
}

func main1() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	//node3 := &ListNode{Val: 3}
	//node4 := &ListNode{Val: 4}
	//node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = nil
	//node2.Next = node3
	//node3.Next = node4
	//node4.Next = node5
	//node5.Next = nil
	reverseKGroup(node1, 2)
}

// 翻转一个链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pNode := head
	cNode := head.Next
	for cNode != nil {
		curNode := cNode.Next
		cNode.Next = pNode
		pNode = cNode
		cNode = curNode
	}
	head.Next = nil
	return pNode
}
