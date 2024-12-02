package core

import "math/rand"

func RandomBy(source []*Spectrum, num int, needReset func(a *List[int], b *Spectrum) bool) *List[int] {
	list := NewList[int]()
	for i := 0; i < num; i++ {
		randomBy(source, list, needReset)
	}
	return list
}

func randomBy(source []*Spectrum, list *List[int], needReset func(a *List[int], b *Spectrum) bool) {
	l := source[rand.Intn(len(source))]
	if needReset(list, l) {
		randomBy(source, list, needReset)
	} else {
		list.Concat(l.List)
	}
}
