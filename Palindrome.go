package palindrome

// Check returns true if the string is a palindrome
func Check(s string) bool {
	var l = len(s)
	for index := 0; index < l/2; index++ {
		if s[index] != s[l-1-index] {
			return false
		}
	}
	return true
}
