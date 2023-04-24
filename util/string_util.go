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
