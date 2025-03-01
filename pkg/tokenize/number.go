package tokenize

import (
	"bytes"
	"errors"
	"io"
)

var (
	ErrUnexpectedDecimal  = errors.New("unexpected decimal character")
	ErrUnexpectedExponent = errors.New("unexpected exponent character")
	ErrUnexpectedSign     = errors.New("unexpected sign character")
)

func consumeLiteralNumber(b []byte, atEOF bool) (int, []byte, error) {
	if len(b) < 1 && atEOF {
		return 0, nil, io.EOF
	} else if len(b) < 1 {
		return 0, nil, nil
	}

	var closedAt int
	var sign bool
	var decimal bool
	var exponent bool

Consuming:
	for i := range b {
		switch b[i] {
		case '-', '+':
			if sign {
				return 0, nil, ErrUnexpectedSign
			}
			sign = true
		case '.':
			if decimal {
				return 0, nil, ErrUnexpectedDecimal
			}
			decimal = true
		case 'e', 'E':
			if exponent {
				return 0, nil, ErrUnexpectedExponent
			}
			exponent = true
			sign, decimal = false, false
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue Consuming
		case 'N':
			switch {
			case i == 0 && bytes.HasPrefix(b, Token{'N', 'a', 'N'}):
				closedAt = 3
			default:
				closedAt = i
			}
			break Consuming
		case 'I':
			switch {
			case i == 0 && bytes.HasPrefix(b, Token{'I', 'n', 'f'}):
				closedAt = 3
			case i == 1 && sign && bytes.HasPrefix(b[1:], Token{'I', 'n', 'f'}):
				closedAt = 4
			default:
				closedAt = i
			}
			break Consuming
		default:
			closedAt = i
			break Consuming
		}
	}

	switch {
	case closedAt == 0 && atEOF:
		return len(b), b, nil
	case closedAt == 0 && !atEOF:
		return 0, nil, nil
	case closedAt == len(b)-1 && atEOF:
		return len(b), b, nil
	case closedAt == len(b)-1 && !atEOF:
		return 0, nil, nil
	default:
		return closedAt, b[:closedAt], nil
	}
}
