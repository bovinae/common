package util

func GetPrefix(prefix, curr []byte) []byte {
	i, j := 0, 0
	for i < len(prefix) && j < len(curr) {
		if prefix[i] != curr[j] {
			break
		}
		i++
		j++
	}
	return curr[:i]
}

func GetSuffix(suffix, curr []byte) []byte {
	i, j := len(suffix)-1, len(curr)-1
	for i >= 0 && j >= 0 {
		if suffix[i] != curr[j] {
			break
		}
		i--
		j--
	}
	return curr[j+1:]
}
