package sort

// quicksort if proper for arrays, while merge can handle list well !!!
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	secHead := Split(head)
	return MergeList(SortList(head), SortList(secHead))
}

// for the division of the list, using fast,slow which means if slow moves the half speed of the fast
func Split(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	tail.Next = nil
	return slow
}

func MergeList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	// merge list consider different between arrays
	var head, cur, p *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur = left
			left = left.Next
		} else {
			cur = right
			right = right.Next
		}
		if head == nil {
			head = cur
		} else {
			p.Next = cur
		}
		p = cur
	}
	if left == nil {
		p.Next = right
	} else {
		p.Next = left
	}
	return head
}
