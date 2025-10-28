package piscine

func IsPositiveNode(node *NodeL) bool {
	if v, ok := node.Data.(int); ok {
		return v > 0
	}
	return false
}

func IsAlNode(node *NodeL) bool {
	_, ok := node.Data.(int)
	return !ok
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l == nil || l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}
