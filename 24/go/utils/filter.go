package utils

import "slices"

func Filter[T any](s []T, fn func(T) bool) []T {
	res := make([]T, 0, len(s))

	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}

	return slices.Clip(res)
}
