package valid_palindrome

func isPalindrome(s string) bool {
	filtered := filter(s)
	i := 0
	j := len(filtered) - 1
	for i < len(filtered)/2 {
		if filtered[i] != filtered[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func filter(s string) string {
	filtered := ""
	for _, v := range s {
		if isAlphaNum(v) {
			if v >= 65 && v <= 90 {
				filtered += string(v + 32)
			} else {
				filtered += string(v)
			}
		}
	}

	return filtered
}

func isAlphaNum(r rune) bool {
	if r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122 {
		return true
	}

	return false
}
