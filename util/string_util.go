package util

func ReverseString(str string) string {
	strLen := len(str)
	if strLen == 0 {
		return ""
	}

	reverse := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		reverse[i] = str[strLen-i-1]
	}

	return string(reverse)
}

func ContainsSubSequence(s, subseq string) bool {
	if len(s) == 0 || len(subseq) == 0 {
		return false
	}

	a := []rune(s)
	b := []rune(subseq)
	var start int
	for i := 0; i < len(b); i++ {
		j := start
		for ; j < len(a); j++ {
			if a[j] == b[i] {
				start = j + 1
				break
			}
		}
		if j >= len(a) {
			return false
		}
	}
	return true
}
