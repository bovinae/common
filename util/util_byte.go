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

func ReverseRuneSlice(rs []rune) []rune {
	rsLen := len(rs)
	if rsLen == 0 {
		return rs
	}

	for i := 0; i < rsLen/2; i++ {
		rs[i], rs[rsLen-i-1] = rs[rsLen-i-1], rs[i]
	}
	return rs
}

func ReverseByteSlice(rs []byte) []byte {
	rsLen := len(rs)
	if rsLen == 0 {
		return rs
	}

	for i := 0; i < rsLen/2; i++ {
		rs[i], rs[rsLen-i-1] = rs[rsLen-i-1], rs[i]
	}
	return rs
}

func RunesCompare(a, b []rune) int {
	k, m := 0, 0
	for k < len(a) && m < len(b) {
		if a[k] < b[m] {
			return -1
		}
		if a[k] > b[m] {
			return 1
		}
		k++
		m++
	}
	if k < len(a) {
		return 1
	}
	if m < len(b) {
		return -1
	}
	return 0
}
