package utils

func Reduce[T any, U any](input []T, initial U, f func(U, T) U) U {
	result := initial
	for _, v := range input {
		result = f(result, v)
	}
	return result
}