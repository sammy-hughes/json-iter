package tokenize

func consumeLiteralComma(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{','}, atEOF)
}
