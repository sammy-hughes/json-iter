package tokenize

import (
	"bufio"
	"io"
	"iter"
)

func Tokens(reader io.Reader) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		scanner := bufio.NewScanner(reader)
		scanner.Split(SplitTokens)

		var ok bool
		for scanner.Scan() {
			if ok = yield(scanner.Bytes()); !ok {
				break
			}
		}
	}
}
