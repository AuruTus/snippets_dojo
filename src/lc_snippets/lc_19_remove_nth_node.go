package lc_snippets

import (
	"context"
	snippets "snippets_dojo/src"
)

type LC19Tstr struct{}

var _ snippets.Tstr = (*LC19Tstr)(nil)

func (t *LC19Tstr) Test(ctx context.Context) error {
	removeNthFromEnd(nil, 1)
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	first, prev := head, &ListNode{Next: head}

	for i := n - 1; i > 0; i-- {
		first = first.Next
	}
	for first.Next != nil {
		prev = prev.Next
		first = first.Next
	}

	if prev.Next == head {
		return head.Next
	}
	prev.Next = prev.Next.Next

	return head
}
