package tokenize

func consumeObjectOpen(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{'{'}, atEOF)
}

func consumeObjectClose(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{'}'}, atEOF)
}
