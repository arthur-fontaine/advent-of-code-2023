package utils

import (
	"strings"
)

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

func RotateString(s string) string {
	rows := strings.Split(s, "\n")
	new_rows := []string{}

	for _, row := range rows {
		for column_index, character := range row {
			if len(new_rows)-1 < column_index {
				new_rows = append(new_rows, string(character))
			} else {
				new_rows[column_index] += string(character)
			}
		}
	}

	return strings.Join(new_rows, "\n")
}
