package logestpalindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
	s := "babad"
	result := longestPalindrome(s)
	if result != "bab" && result != "aba" {
		t.Errorf("Expected bab or aba but got %s", result)
	}
}
