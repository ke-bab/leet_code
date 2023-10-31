package reverse_k_group

import "testing"

func TestPlusOne(t *testing.T) {

	testCases := []struct {
		input *ListNode
		k     int
		want  *ListNode
	}{
		{
			input: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}},
			k:     3,
			want:  &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}},
		},
	}

	for _, tc := range testCases {
		got := reverseKGroup(tc.input, tc.k)
		if !ok(got, tc.want) {
			t.Fatalf("result is wrong")
		}
	}
}

func ok(got *ListNode, wanted *ListNode) bool {
	current := got
	currentWanted := wanted
	for {
		if current.Val != currentWanted.Val {
			return false
		}

		if current.Next != nil && currentWanted.Next == nil {
			return false
		}
		if current.Next == nil && currentWanted.Next != nil {
			return false
		}

		if current.Next != nil && currentWanted.Next != nil {
			current = current.Next
			currentWanted = currentWanted.Next
		} else {
			break
		}
	}

	return true
}
