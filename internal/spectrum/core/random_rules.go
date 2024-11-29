package core

func ResetRules(rules ...func(a, b *List[int]) bool) func(a, b *List[int]) bool {
	return func(a, b *List[int]) bool {
		for _, rule := range rules {
			if rule(a, b) {
				return true
			}
		}
		return false
	}
}

func RuleSameFoot(a, b *List[int]) bool {
	if a.Head == nil || b.Head == nil {
		return false
	}
	if a.Tail.Data == b.Head.Data {
		return true
	}
	return false
}

func RuleReverse(a, b *List[int]) bool {
	if a.Head == nil || b.Head == nil {
		return false
	}
	if a.Tail.Prev != nil && a.Tail.Prev.Data == 0 && a.Tail.Data == 3 && b.Head.Data == 4 {
		return true
	}
	if a.Tail.Prev != nil && a.Tail.Prev.Data == 3 && a.Tail.Data == 0 && b.Head.Data == 1 {
		return true
	}
	return false
}

func RuleDiagonal(a, b *List[int]) bool {
	if a.Head == nil || b.Head == nil {
		return false
	}
	if a.Tail.Prev != nil && a.Tail.Prev.Data == 0 && a.Tail.Data == 2 && b.Head.Data == 4 {
		return true
	}
	if a.Tail.Prev != nil && a.Tail.Prev.Data == 3 && a.Tail.Data == 2 && b.Head.Data == 1 {
		return true
	}
	return false
}

func RuleRepeat(a, b *List[int]) bool {
	if a.Head == nil || b.Head == nil {
		return false
	}
	if a.Tail.Prev != nil && b.Head != nil && a.Tail.Prev.Data == b.Head.Data {
		return true
	}
	if a.Tail != nil && b.Head.Next != nil && a.Tail.Data == b.Head.Next.Data {
		return true
	}
	return false
}

func RuleNoRepeat(a, b *List[int]) bool {
	if a.Head == nil || b.Head == nil {
		return false
	}
	return !RuleRepeat(a, b)
}
