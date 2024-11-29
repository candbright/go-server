package utils

import "reflect"

func Contains[T any](arr []T, value T) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}
