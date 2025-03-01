package tokenize

import (
	"bytes"
	"errors"
	"io"
)

var ErrUnmatched = errors.New("token not matched")

func consumeLiteralConstant(b, token []byte, atEOF bool) (int, []byte, error) {
	switch {
	case len(b) < 1:
		return 0, nil, nil
	case b[0] != token[0]:
		return 0, nil, ErrUnmatched
	case len(b) < len(token) && atEOF:
		return 0, nil, nil
	case len(b) < len(token):
		return 0, nil, io.ErrUnexpectedEOF
	case bytes.HasPrefix(b, token):
		return len(token), token, nil
	default:
		return 0, nil, ErrUnmatched
	}
}
