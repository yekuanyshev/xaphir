package utils

func SliceMap[X any, Y any](s []X, m func(X) Y) []Y {
	mapped := make([]Y, len(s))

	for i, item := range s {
		mapped[i] = m(item)
	}

	return mapped
}

func SliceFilter[X any](s []X, f func(X) bool) []X {
	filtered := make([]X, 0, len(s))

	for _, item := range s {
		if f(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
