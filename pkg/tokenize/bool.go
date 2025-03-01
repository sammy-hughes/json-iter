package tokenize

func consumeLiteralFalse(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{'f', 'a', 'l', 's', 'e'}, atEOF)
}

func consumeLiteralTrue(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{'t', 'r', 'u', 'e'}, atEOF)
}
