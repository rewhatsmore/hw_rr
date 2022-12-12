package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var (
		b   strings.Builder
		n   int
		str = []rune(s)
	)

	for i := 0; i < len(str); i++ {
		r := str[i]
		if unicode.IsDigit(r) {
			return "", ErrInvalidString
		}

		if r == '\\' {
			if unicode.IsDigit(str[i+1]) || str[i+1] == '\\' {
				r = str[i+1]
				i++
			} else {
				return "", ErrInvalidString
			}
		}

		if i != len(str)-1 && unicode.IsDigit(str[i+1]) {
			n = int(str[i+1] - 48)
			i++
		} else {
			n = 1
		}

		b.WriteString(strings.Repeat(string(r), n))
	}

	return b.String(), nil
}
