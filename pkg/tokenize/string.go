package tokenize

import (
	"bytes"
	"errors"
)

var ErrMissingQuote = errors.New("expected matched quotes")

func consumeLiteralString(b []byte, atEOF bool) (int, []byte, error) {
	if !bytes.HasPrefix(b, []byte{'"'}) {
		return 0, nil, ErrMissingQuote
	} else if len(b) < 2 {
		return 0, nil, ErrMissingQuote
	}

	var closeAt int
	var escape bool
ScanningLoop:
	for i := range len(b) - 1 {
		switch {
		case b[i+1] == '"' && !escape:
			closeAt = i + 2
			break ScanningLoop
		case b[i+1] == '\\':
			escape = !escape
		default:
			continue ScanningLoop
		}
	}

	switch {
	case closeAt == 0 && atEOF:
		return 0, nil, ErrMissingQuote
	case closeAt == 0:
		return 0, nil, nil
	case closeAt == len(b):
		return closeAt, b, nil
	default:
		return closeAt, b[:closeAt], nil
	}
}
