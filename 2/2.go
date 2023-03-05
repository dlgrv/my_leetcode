package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// l1 := ListNode{0, nil}
	// l2 := ListNode{0, nil}
	// l2 := ListNode{0, &ListNode{8, &ListNode{6, &ListNode{5, &ListNode{6, &ListNode{8, &ListNode{3, &ListNode{5, &ListNode{7, nil}}}}}}}}}
	// l1 := ListNode{6, &ListNode{7, &ListNode{8, &ListNode{0, &ListNode{8, &ListNode{5, &ListNode{8, &ListNode{9, &ListNode{7, nil}}}}}}}}}
	// l1 := ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}}}
	// l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// l1 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	// l2 := ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	fmt.Println(addTwoNumbers(&l1, &l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: 0}
	tmp := res
	sum := 0

	for !(l1 == nil && l2 == nil && sum == 0) {
		val1 := 0
		val2 := 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum += val1 + val2

		tmp.Next = &ListNode{sum % 10, nil}
		tmp = tmp.Next

		sum = sum / 10
	}

	return res.Next
}

// func listNodeIntoString(numbs string, ln *ListNode) string {
// 	if ln.Next == nil {
// 		numbs += strconv.Itoa(ln.Val)
// 		return numbs
// 	} else {
// 		numbs += strconv.Itoa(ln.Val)
// 		return listNodeIntoString(numbs, ln.Next)
// 	}
// }

// func reverse(s string) (res string) {
// 	for _, v := range s {
// 		res = string(v) + res
// 	}
// 	return res
// }

// func integerIntoListNode(numbs string) *ListNode {
// 	if numbs == "" {
// 		return nil
// 	} else {
// 		addNumbs, _ := strconv.Atoi(numbs[len(numbs)-1:])
// 		remainder := numbs[:len(numbs)-1]
// 		return &ListNode{addNumbs, integerIntoListNode(remainder)}
// 	}
// }
