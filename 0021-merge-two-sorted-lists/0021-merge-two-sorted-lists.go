/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var result []int

	for true {
		// if list1.Next == nil && len(result) == 0 {
		if list1 == nil {
			break
		}
		i := 0
		for true {
			if len(result) == 0 {
				result = append(result, list1.Val)
				break
			} else {
				if len(result) == i {
					result = append(result, list1.Val)
					break
				}
				if result[i] > list1.Val {
					result = append(result[:i+1], result[i:]...)
					result[i] = list1.Val
					break
				}
			}
			i++
		}
		if list1.Next == nil {
			break
		}
		list1 = list1.Next
	}

	for true {
		// if list2.Next == nil && len(result) == 0 {
		if list2 == nil {
			break
		}
		i := 0
		for true {
			if len(result) == 0 {
				result = append(result, list2.Val)
				break
			} else {
				if len(result) == i {
					result = append(result, list2.Val)
					break
				}
				if result[i] > list2.Val {
					result = append(result[:i+1], result[i:]...)
					result[i] = list2.Val
					break
				}
			}
			i++
		}
		if list2.Next == nil {
			break
		}
		list2 = list2.Next
	}

	if len(result) == 0 {
		return nil
	} else {
		node := &ListNode{Val: result[len(result)-1]}
		for i := len(result) - 2; i >= 0; i-- {
			newNode := &ListNode{Val: result[i], Next: node}
			node = newNode
		}
		return node
	}
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

// 	var result []int

// 	for list1 != nil {
// 		result = append(result, list1.Val)
// 		list1 = list1.Next
// 	}
// 	for list2 != nil {
// 		result = append(result, list2.Val)
// 		list2 = list2.Next
// 	}
// 	sort.Ints(result)

// 	dummy := &ListNode{}
// 	current := dummy
// 	for _, val := range result {
// 		newNode := &ListNode{Val: val}
// 		current.Next = newNode
// 		current = newNode
// 	}

// 	return dummy.Next
// }