// Time: O(n)
// Space: O(n)
func hasCycle(head *ListNode) bool {
	// Space: O(n)
	current, set := head, NewSet()

	// Time: O(n)
	for current != nil {
		if set.Has(current) {
			return true
		}
		set.Add(current)
		current = current.Next
	}
	return false
}

type Set struct {
	items map[*ListNode]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[*ListNode]struct{}),
	}
}

func (s *Set) Has(n *ListNode) bool {
	_, exists := s.items[n]
	return exists
}

func (s *Set) Add(n *ListNode) {
	s.items[n] = struct{}{}
}