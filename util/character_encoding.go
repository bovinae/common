package util

import (
	"bytes"
	"io"

	"github.com/bovinae/golang.org/x/text/transform"
	"github.com/pkg/errors"
)

// transform UTF8 rune into GBK
func UTF82GBK(reader *transform.Reader, src string) ([]byte, error) {
	reader.Reset(bytes.NewReader([]byte(src)))
	gbkBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "io.ReadAll")
	}
	return gbkBytes, nil
}

// GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(reader *transform.Reader, src []byte) (string, error) {
	reader.Reset(bytes.NewReader(src))
	utf8Bytes, err := io.ReadAll(reader)
	if err != nil {
		return "", errors.Wrap(err, "io.ReadAll")
	}
	return string(utf8Bytes), nil
}
