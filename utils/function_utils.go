package utils

func Memoize[V comparable, T any](f func(v V) T) func(v V) T {
	cache := map[V]T{}

	return func(v V) T {
		if r, ok := cache[v]; ok {
			return r
		}

		r := f(v)

		cache[v] = r

		return r
	}
}
