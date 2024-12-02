package core

import "github.com/candbright/go-server/internal/base/utils"

func ResetRules(rules ...func(a, b Spectrum) bool) func(a, b Spectrum) bool {
	return func(a, b Spectrum) bool {
		for _, rule := range rules {
			if rule(a, b) {
				return true
			}
		}
		return false
	}
}

func RuleSameFoot(a, b Spectrum) bool {
	if a.LastFoot() == b.FirstFoot() {
		return true
	}
	return false
}

func RuleSameNote(a, b Spectrum) bool {
	return !RuleSameFoot(a, b) && utils.ArraysEqual(a.LastNote().Position, b.FirstNote().Position)
}

func RuleReverse(a, b Spectrum) bool {
	if !a.LastNote().Single() || !b.FirstNote().Single() {
		return false
	}
	if a.Tail().Prev != nil && a.Tail().Prev.Data.Position.At(0) && a.Tail().Data.Position.At(3) && b.Head().Data.Position.At(4) {
		return true
	}
	if a.Tail().Prev != nil && a.Tail().Prev.Data.Position.At(3) && a.Tail().Data.Position.At(0) && b.Head().Data.Position.At(1) {
		return true
	}
	return false
}

func RuleDiagonal(a, b Spectrum) bool {
	if !a.LastNote().Single() || !b.FirstNote().Single() {
		return false
	}
	if a.Tail().Prev != nil && a.Tail().Prev.Data.Position.At(0) && a.Tail().Data.Position.At(2) && b.Head().Data.Position.At(4) {
		return true
	}
	if a.Tail().Prev != nil && a.Tail().Prev.Data.Position.At(3) && a.Tail().Data.Position.At(2) && b.Head().Data.Position.At(1) {
		return true
	}
	return false
}

func RuleRepeat(a, b Spectrum) bool {
	return utils.ArraysEqual(a.Tail().Data.Position, b.Head().Data.Position)
}

func RuleNoRepeat(a, b Spectrum) bool {
	return !RuleRepeat(a, b)
}
