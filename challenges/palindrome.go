package challenges

func IsPalindrome(text string) bool {

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-(1+i)] {
			return false
		}
	}

	return true
}
