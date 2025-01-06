package utils

func Map[T, V any](s []T, fn func(T) V) []V {
	res := make([]V, len(s))
	for i, v := range s {
		res[i] = fn(v)
	}
	return res
}
