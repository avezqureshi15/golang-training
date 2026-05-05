package utils

func Filter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}