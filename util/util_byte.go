package util

func GetPrefix[T byte | rune](prefix, curr []T) []T {
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

func GetSuffix[T byte | rune](suffix, curr []T) []T {
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

func ReverseSlice[T byte | rune](rs []T) []T {
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
