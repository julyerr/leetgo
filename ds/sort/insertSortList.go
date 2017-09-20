package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解题思路，将整个list划分为两部分，第一部分作为已经排序好的，第二部分作为等待插入排序的
// 将不符合顺序的node取出来，插入到第一部分
func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	HeadNode := &ListNode{Next: head}
	for cur := head; cur.Next != nil; {
		p := cur.Next
		if cur.Val > p.Val {
			cur.Next = p.Next
			pre, start := HeadNode, HeadNode.Next
			for ; start != cur && start.Val < p.Val; start = start.Next {
				pre = start
			}
			pre.Next = p
			p.Next = start
		} else {
			cur = p
		}
	}
}
