package tokenize

import (
	"strconv"
)

type Token []byte

func (token Token) String() string {
	return strconv.Quote(string(token))
}

func (token Token) Equals(other Token) bool {
	if len(token) != len(other) {
		return false
	}

	for i := range token {
		if token[i] != other[i] {
			return false
		}
	}

	return true
}
