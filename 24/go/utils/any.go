package utils

func Any[T any](s []T, fn func(v T) bool) bool {

	for _, v := range s {
		if fn(v) {
			return true
		}
	}

	return false
}
