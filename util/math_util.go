package util

func Min(a, b any) any {
	if CompareAny(a, b) == LESS {
		return a
	}
	return b
}

func Max(a, b any) any {
	if CompareAny(a, b) == GREATER {
		return a
	}
	return b
}
