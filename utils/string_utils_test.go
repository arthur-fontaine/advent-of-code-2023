package utils_test

import (
	"arthur-fontaine/advent-of-code-2023/utils"
	"testing"
)

func TestRotateString(t *testing.T) {
	s := "ABC\nDEF"
	want := "AD\nBE\nCF"
	rotated_string := utils.RotateString(s)
	if want != rotated_string {
		t.Fatalf(`RotateString("ABC\\nDEF") = %q, want match for %#q`, rotated_string, want)
	}
}
