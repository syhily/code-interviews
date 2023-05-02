package chapter12

import (
	"github.com/syhily/code-interviews/common"
)

func mergeSort(numbers ...int) []int {
	length, res, merge := len(numbers), numbers, make([]int, len(numbers))

	for step := 1; step < length; step *= 2 {
		for start := 0; start < length; start += step * 2 {
			mid := common.MinNumber(start+step, length)
			end := common.MinNumber(mid+step, length)
			l, r, i := start, mid, start

			for l < mid || r < end {
				if r == end || (l < mid && res[l] < res[r]) {
					merge[i] = res[l]
					l++
				} else {
					merge[i] = res[r]
					r++
				}
				i++
			}
		}

		res, merge = merge, res
	}

	return res
}

// O(log N) time, O(log N) memory.
func reclusiveMergeSortList(node *common.ListNode) *common.ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	head1, head2 := node, splitList(node)

	head1 = reclusiveMergeSortList(head1)
	head2 = reclusiveMergeSortList(head2)

	return mergeList(head1, head2)
}

func splitList(node *common.ListNode) *common.ListNode {
	slow, fast := node, node.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow, slow.Next = slow.Next, nil

	return slow
}

func mergeList(a, b *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{}

	curr := dummy
	for a != nil && b != nil {
		if a.Value < b.Value {
			curr.Next = a
			a = a.Next
		} else {
			curr.Next = b
			b = b.Next
		}

		curr = curr.Next
	}
	if a == nil {
		curr.Next = b
	} else {
		curr.Next = a
	}

	return dummy.Next
}

// O(log N) time, O(1) memory.
func mergeSortList(node *common.ListNode) *common.ListNode {
	length := 0
	for n := node; n != nil; n = n.Next {
		length++
	}

	dummy, step := &common.ListNode{Next: node}, 1
	for step < length {
		start := dummy

		for i := 0; i < length; i += step * 2 {
			// Find the middle node.
			middle := start
			for k := 0; k < step && middle != nil; k++ {
				middle = middle.Next
			}

			// Merge two sorted lists.
			start.Next, start = mergeSubList(start.Next, middle.Next, step)
		}

		step *= 2
	}

	return dummy.Next
}

// Merge the two lists and return the head node.
func mergeSubList(a, b *common.ListNode, step int) (*common.ListNode, *common.ListNode) {
	dummy, al, bl := &common.ListNode{}, 0, 0

	curr := dummy
	for al < step && bl < step && a != nil && b != nil {
		if a.Value < b.Value {
			curr.Next, a = a, a.Next
			al++
		} else {
			curr.Next, b = b, b.Next
			bl++
		}

		curr = curr.Next
	}

	for al < step {
		curr.Next, curr, a = a, a, a.Next
		al++
	}
	for bl < step && b != nil {
		curr.Next, curr, b = b, b, b.Next
		bl++
	}

	curr.Next = b

	return dummy.Next, curr
}
