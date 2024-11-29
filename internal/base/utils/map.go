package utils

func ToArray[V any](m map[string]V) []V {
	if len(m) == 0 {
		return make([]V, 0)
	}
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
