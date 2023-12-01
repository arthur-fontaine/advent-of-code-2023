package utils

func StringEndsWith(s string, with string) bool {
	return len(s) >= len(with) && s[len(s)-len(with):] == with
}

func Reverse(s string) string {
	reversed_string := ""
	for _, character := range s {
		reversed_string = string(character) + reversed_string
	}
	return reversed_string
}
