package utils

func AllInArray[V comparable](array []V, value V) bool {
	for _, v := range array {
		if v != value {
			return false
		}
	}

	return true
}

func ArrayIncludes[V comparable](array []V, value V) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func MapArray[V comparable, T comparable](array []V, callback func(v V) (T, error)) ([]T, error) {
	new_array := make([]T, len(array))

	for i, v := range array {
		new_v, err := callback(v)
		if err != nil {
			return new_array, err
		}
		new_array[i] = new_v
	}

	return new_array, nil
}

func ArraysAreSame[V comparable](array_a []V, array_b []V) bool {
	if len(array_a) != len(array_b) {
		return false
	}

	for i := range array_a {
		if array_a[i] != array_b[i] {
			return false
		}
	}

	return true
}
