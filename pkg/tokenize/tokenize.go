package tokenize

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var ErrUnmatched = errors.New("token not matched")

type Token []byte

func scanLiteralConstant(b, token []byte) (int, []byte, error) {
	switch {
	case len(b) < 1:
		return 0, nil, nil
	case b[0] != token[0]:
		return 0, nil, ErrUnmatched
	case len(b) < len(token):
		return 0, nil, nil
	case bytes.HasPrefix(b, token):
		return len(token), token, nil
	default:
		return 0, nil, ErrUnmatched
	}
}

func scanLiteralNumber(b []byte, atEOF bool) (int, []byte, error) {
	if len(b) < 1 && atEOF {
		return 0, nil, io.EOF
	} else if len(b) < 1 {
		return 0, nil, nil
	}

	isNumeral := func(r rune) bool { return r >= '0' && r <= '9' }
	isSigned := func(r rune) bool { return r == '-' || r == '+' }
	isIEEE754E := func(r rune) bool { return r == 'e' || r == 'E' }
	isDecimal := func(r rune) bool { return r == '.' || isNumeral(r) || isSigned(r) }
	isNonNumeric := func(r rune) bool { return !(isDecimal(r) || isIEEE754E(r)) }

	length := bytes.IndexFunc(b, isNonNumeric)
	if length == -1 && !atEOF {
		return 0, nil, nil
	} else if length == -1 {
		length = len(b)
	}

	if bytes.IndexFunc(b[:length], isIEEE754E) != -1 {
		ieee754 := bytes.FieldsFunc(b[:length], isIEEE754E)
		if len(ieee754) > 2 {
			return 0, nil, ErrUnmatched
		} else if bytes.IndexFunc(ieee754[0], isNonNumeric) != -1 {
			return 0, nil, ErrUnmatched
		} else if bytes.IndexFunc(ieee754[1], isNonNumeric) != -1 {
			return 0, nil, ErrUnmatched
		}
	}

	return length, b[:length], nil
}

func SplitTokens(b []byte, atEOF bool) (int, []byte, error) {
	if len(b) < 1 && atEOF {
		return 0, nil, io.EOF
	} else if len(b) < 1 {
		return 0, nil, nil
	}

	switch b[0] {
	case '"':

	case 'n':
		return scanLiteralConstant(b, Token{'n', 'u', 'l', 'l'})
	case 't':
		return scanLiteralConstant(b, Token{'t', 'r', 'u', 'e'})
	case 'f':
		return scanLiteralConstant(b, Token{'f', 'a', 'l', 's', 'e'})
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '+':
		return scanLiteralNumber(b, atEOF)
	case ',':
		return scanLiteralConstant(b, Token{','})
	case '[':
		return scanLiteralConstant(b, Token{'['})
	case ']':
		return scanLiteralConstant(b, Token{']'})
	case '{':
		return scanLiteralConstant(b, Token{'{'})
	case '}':
		return scanLiteralConstant(b, Token{'}'})
	}

	return 0, nil, fmt.Errorf("%w; token=%s", ErrUnmatched, string(b))
}
