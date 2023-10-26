package reverse__k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// start from first pos
	// check if right block is k group
	n, ok := getEndOfNextGroup(head, k)
	if ok {
		temp := head.Next
		head.Next = n

	} else {

	}
	// is yes - swap current and last pos of right block
	// if not - swap current with first pos of right block

	// remember address of old node (because we changed it)
	// go to next pos in block
	// do the same but remember next pos address and swap it with previous pos address
	// step next and repeat until we find end of block

}

func thisGroupIsKGroup(start *ListNode, k int) bool {
	c := 0
	for e := start; e != nil; e = e.Next {
		c++
		if c == k-1 {
			return true
		}
	}

	return false
}

func getEndOfNextGroup(start *ListNode, k int) (*ListNode, bool) {
	c := 0
	var lastNode *ListNode
	for e := start; e != nil; e = e.Next {
		c++
		lastNode = e
		if c == (k-1)*2 {
			return e, true
		}
	}

	return lastNode, false
}
