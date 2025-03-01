package tokenize

func consumeArrayOpen(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{'['}, atEOF)
}

func consumeArrayClose(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{']'}, atEOF)
}

func consumeLiteralColon(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{':'}, atEOF)
}
