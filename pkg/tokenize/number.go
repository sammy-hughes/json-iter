package tokenize

import (
	"bytes"
	"io"
)

func consumeLiteralNumber(b []byte, atEOF bool) (int, []byte, error) {
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
