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
