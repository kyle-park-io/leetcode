/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var result []int

	for list1 != nil {
		result = append(result, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		result = append(result, list2.Val)
		list2 = list2.Next
	}
	sort.Ints(result)

	dummy := &ListNode{}
	current := dummy
	for _, val := range result {
		newNode := &ListNode{Val: val}
		current.Next = newNode
		current = newNode
	}

	return dummy.Next
}