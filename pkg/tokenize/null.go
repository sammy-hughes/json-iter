package tokenize

func consumeLiteralNull(b []byte, atEOF bool) (int, []byte, error) {
	return consumeLiteralConstant(b, Token{'n', 'u', 'l', 'l'}, atEOF)
}
