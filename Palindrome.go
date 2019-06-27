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

// MakeFrom returns a palindrome from the string in input.
// s is the input string
// even is a flag that determines if the return string is of even size or odd, which in turn is determined by duplicating or not the last char of s.
func MakeFrom(s string, even bool) string {
	var l = len(s)
	if l == 0 {
		return s
	}
	if l == 1 {
		if even {
			return s + s
		}
		return s
	}
	var buf = 1
	if even {
		buf = 0
	}
	var res = s
	for index := buf; index < l; index++ {
		var c = string(s[l-index-1])
		res = res + c
	}
	return res
}
