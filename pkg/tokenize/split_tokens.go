package tokenize

import (
	"fmt"
	"io"
)

func SplitTokens(b []byte, atEOF bool) (int, []byte, error) {
	if len(b) < 1 && atEOF {
		return 0, nil, io.EOF
	} else if len(b) < 1 {
		return 0, nil, nil
	}

	switch b[0] {
	case 'n':
		return consumeLiteralNull(b, atEOF)
	case 'f':
		return consumeLiteralFalse(b, atEOF)
	case 't':
		return consumeLiteralTrue(b, atEOF)
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '+':
		return consumeLiteralNumber(b, atEOF)
	case '"':
		return consumeLiteralString(b, atEOF)
	case ',':
		return consumeLiteralComma(b, atEOF)
	case '[':
		return consumeArrayOpen(b, atEOF)
	case ']':
		return consumeArrayClose(b, atEOF)
	case '{':
		return consumeObjectOpen(b, atEOF)
	case '}':
		return consumeObjectClose(b, atEOF)
	case ':':
		return consumeLiteralColon(b, atEOF)
	case '\n', '\r', '\t', ' ':
		n, b, atEOF := SplitTokens(b[1:], atEOF)
		return n + 1, b, atEOF
	}

	return 0, nil, fmt.Errorf("%w; token=%s", ErrUnmatched, string(b))
}
