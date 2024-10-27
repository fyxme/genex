package genex

import (
	"regexp/syntax"
)

type GenOpts struct {
	charset *syntax.Regexp
	infinite int
}

const (
	DEFAULT_CHARSET = "[a-zA-Z0-9]"
	DEFAULT_INFINITE = 3
)

func valuesFromOptions(options *GenOpts) (*syntax.Regexp, int) {
	infinite := DEFAULT_INFINITE
	var charset, _ = syntax.Parse(DEFAULT_CHARSET, syntax.Perl)

	if options != nil {
		if options.charset != nil {
			charset = options.charset
		}
		infinite = options.infinite
	}

	return charset, infinite
}

