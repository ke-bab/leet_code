package reverse_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return reverseGroup(head, k)
}

func canBeReversed(head *ListNode, k int) bool {
	current := head
	for i := 0; i < k; i++ {
		if current == nil {
			return false
		}
		current = current.Next
	}

	return true
}

func reverseGroup(head *ListNode, k int) *ListNode {
	var first *ListNode
	var last *ListNode
	var prev *ListNode
	var current *ListNode
	var next *ListNode

	last = head
	current = head

	for i := 0; i < k; i++ {
		next = current.Next
		current.Next = prev
		// iteration step
		prev = current
		current = next
	}
	first = prev
	last.Next = current // link "new" last node and next group node
	if canBeReversed(last.Next, k) {
		first := reverseGroup(last.Next, k)
		last.Next = first
	}

	return first
}
