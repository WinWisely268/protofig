package utils

import (
	"fmt"
	"unicode/utf8"
)

// Parts of the code are taken/inspired from https://github.com/burntsushi/toml
// Give credits where it's due.

/*
Legend:
	r is rune
*/

type tknType int

const (
	tknError tknType = iota
	tknNIL           // used in the parser to indicate no type
	tknEOF
	tknString
	tknMultilineString
	tknCompStart
	tknCompEnd
	tknKeyStart
	tknKeyEnd
	tknCommentStart
	tknText
)

const (
	eof           = 0
	compNameStart = '['
	compNameEnd   = ']'
	keyStart      = ':'
	keyEnd        = ':'
	commentStart  = '#'
	stringStart   = '`'
	stringEnd     = '`'
)

// recursive function
type stateFn func(lx *lexer) stateFn

type lexer struct {
	input      string
	start      int
	pos        int
	line       int
	state      stateFn
	tkns       chan tkn
	prevWidths [2]int
	nprev      int // how many of prevWidths are in use
	atEOF      bool
	// A stack of state functions used to maintain context.
	// The idea is to reuse parts of the state machine in various places.
	// For example, values can appear at the top level or within arbitrarily
	// nested arrays. The last state on the stack is used after a value has
	// been lexed. Similarly for comments.
	stack []stateFn
}

type tkn struct {
	typ  tknType
	val  string
	line int
}

func (lx *lexer) nextToken() tkn {
	// consume token chan
	for {
		select {
		case tkn := <-lx.tkns:
			return tkn
		default:
			lx.state = lx.state(lx)
		}
	}
}

func newlex(input string) *lexer {
	return &lexer{
		input: input,
		start: 0,
		line:  1,
		tkns:  make(chan tkn, 10),
		stack: make([]stateFn, 0, 10),
	}
}

// next returns rune next to current pos
func (lx *lexer) next() (r rune) {
	if lx.atEOF {
		panic(fmt.Sprintf(errCalledAfterEOF, "next"))
	}
	if lx.pos >= len(lx.input) {
		lx.atEOF = true
		return eof
	}
	if lx.input[lx.pos] == '\n' {
		lx.line++
	}
	lx.prevWidths[1] = lx.prevWidths[0]
	if lx.nprev < 2 {
		lx.nprev++
	}
	r, w := utf8.DecodeRuneInString(lx.input[lx.pos:])
	lx.prevWidths[0] = w
	lx.pos += w
	return r
}

// consumes elements at the start / top level of config file
func lexStart(lx *lexer) stateFn {
	r := lx.next()
}
