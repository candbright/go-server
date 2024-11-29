package core

import "math/rand"

func RandomBy(sourceMap map[int]*List[int], num int, needReset func(a, b *List[int]) bool) *List[int] {
	list := NewList[int]()
	for i := 0; i < num; i++ {
		randomBy(sourceMap, list, needReset)
	}
	return list
}

func randomBy(sourceMap map[int]*List[int], list *List[int], needReset func(a, b *List[int]) bool) {
	l := sourceMap[rand.Intn(len(sourceMap))]
	if needReset(list, l) {
		randomBy(sourceMap, list, needReset)
	} else {
		list.Concat(l)
	}
}
