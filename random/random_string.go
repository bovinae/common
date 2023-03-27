// reference https://github.com/kpbird/golang_random_string
package random

import "math/rand"

const (
	_charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	_charIdxBits = 6                   // 6 bits to represent a character index: [0, len(CHAR_SET))
	_charIdxMask = 1<<_charIdxBits - 1 // All 1-bits, as many as charIdxBits
	_charIdxMax  = 63 / _charIdxBits   // # of letter indices fitting in 63 bits
)

func GetRandomString(n int) string {
	b := make([]byte, n)
	l := len(_charSet)
	// A rand.Int63() generates 63 random bits, enough for charIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), _charIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), _charIdxMax
		}
		if idx := int(cache & _charIdxMask); idx < l {
			b[i] = _charSet[idx]
			i--
		}
		cache >>= _charIdxBits
		remain--
	}

	return string(b)
}
