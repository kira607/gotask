/*
Copyright © 2026 kira607 <kirill.lesckin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package task

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Tokenize(input string) ([]Token, error) {
	lex := lexer{}
	return lex.tokenize(input)
}

type lexer struct {
	input  string
	pos    int
	tokens []Token
}

func (l *lexer) peek() (rune, int) {
	return utf8.DecodeRuneInString(l.input[l.pos:])
}

func (l *lexer) peekNext() (rune, int) {
	_, size := l.peek()
	return utf8.DecodeRuneInString(l.input[l.pos+size:])
}

func (l *lexer) peekAt(pos int) (rune, int) {
	return utf8.DecodeRuneInString(l.input[pos:])
}

func (l *lexer) advance(size int) {
	l.pos += size
}

func (l *lexer) atEnd() bool {
	return l.pos >= len(l.input)
}

func (l *lexer) tokenize(input string) ([]Token, error) {
	l.pos = 0
	l.tokens = []Token{}
	l.input = input

	if !utf8.ValidString(input) {
		for i := 0; i < len(input); {
			r, size := utf8.DecodeRuneInString(input[i:])
			if r == utf8.RuneError && size == 1 {
				return nil, l.errorAt(i, "invalid UTF-8")
			}
			i += size
		}
	}

	for !l.atEnd() {
		r, size := l.peek()
		switch {
		case unicode.IsSpace(r): // ' ', '\n', '\t', ...
			l.advance(size)
		case r == '(':
			l.addToken(LEFT_PAREN, "(", size)
		case r == ')':
			l.addToken(RIGHT_PAREN, ")", size)
		case r == '=':
			l.addToken(EQUALS, "=", size)
		case r == '>': // >=
			l.handleTwoCharOperator(">", GREATER_THAN, ">=", GREATER_OR_EQUAL)
		case r == '<': // <=
			l.handleTwoCharOperator("<", LESS_THAN, "<=", LESS_OR_EQUAL)
		case r == '~':
			_, nextSize := l.peekNext()
			if nextSize == 0 || l.input[l.pos+size:l.pos+size+nextSize] != "=" {
				return nil, l.errorAt(l.pos, "expected '=' after '~'")
			}
			l.addToken(FUZZY_MATCH, "~=", size+nextSize)
		case r == '!': // !=
			l.handleTwoCharOperator("!", NOT, "!=", NOT_EQUALS)
		case r == '"':
			e := l.readString()
			if e != nil {
				return nil, e
			}
		default:
			e := l.readWord()
			if e != nil {
				return nil, e
			}
		}
	}

	return l.tokens, nil
}

// Add a token in tokens from current pos with the given size, then advance given size
func (l *lexer) addToken(tokenType TokenType, value string, size int) {
	l.tokens = append(l.tokens, Token{
		Type:  tokenType,
		Value: value,
		Start: l.pos,
		End:   l.pos + size,
	})
	l.advance(size)
}

func (l *lexer) handleTwoCharOperator(op string, opType TokenType, twoOp string, twoOpType TokenType) {
	_, size := l.peek()
	next, nextSize := l.peekNext()

	if nextSize == 0 {
		l.addToken(opType, op, size)
		return
	}

	if next == '=' {
		l.addToken(twoOpType, twoOp, size+nextSize)
		return
	}

	l.addToken(opType, op, size)
}

func (l *lexer) readString() error {
	_, size := l.peek() // should be (34, 1) for '"' (34) which is 1 byte
	start := l.pos
	i := start + size
	for i < len(l.input) {
		r, size := l.peekAt(i)
		if r == '\\' {
			_, escapedSize := l.peekAt(i + size)
			if escapedSize == 0 {
				return l.errorAt(start, "unterminated string literal")
			}
			i += size + escapedSize
			continue
		}
		if r == '"' {
			l.addToken(STRING, l.input[start:i+size], i+size-start)
			return nil
		}
		i += size
	}

	if i >= len(l.input) {
		return l.errorAt(start, "unterminated string literal")
	}
	return nil
}

func (l *lexer) readWord() error {
	start := l.pos
	i := l.pos
	for i < len(l.input) {
		r, size := l.peekAt(i)

		if isDelimiter(r) || size == 0 {
			break
		}

		i += size
	}

	word := l.input[start:i]

	var tokenType TokenType

	switch strings.ToLower(word) {
	case "and":
		tokenType = AND
	case "or":
		tokenType = OR
	case "not":
		tokenType = NOT
	case "has":
		tokenType = HAS
	default:
		tokenType = IDENTIFIER
	}

	l.addToken(tokenType, word, i-start)
	return nil
}

func isDelimiter(r rune) bool {
	return unicode.IsSpace(r) ||
		r == '(' ||
		r == ')' ||
		r == '"' ||
		isOperator(r)
}

func isOperator(r rune) bool {
	return r == '=' ||
		r == '>' ||
		r == '<' ||
		r == '~' ||
		r == '!'
}

func (l *lexer) errorAt(position int, message string) TokenizeErr {
	return TokenizeErr{Input: l.input, Position: position, Message: message}
}

// TokenizeErr reports a lexical error at a byte offset in the original input.
// Its Error method renders the source line, a one-based line and column, and
// a caret pointing at the error location.
type TokenizeErr struct {
	Input    string
	Position int
	Message  string
}

func (e TokenizeErr) Error() string {
	position := min(max(e.Position, 0), len(e.Input))
	lineStart := strings.LastIndex(e.Input[:position], "\n") + 1
	lineEnd := len(e.Input)
	if end := strings.IndexByte(e.Input[position:], '\n'); end >= 0 {
		lineEnd = position + end
	}

	line := e.Input[lineStart:lineEnd]
	lineNumber := strings.Count(e.Input[:lineStart], "\n") + 1
	column := utf8.RuneCountInString(e.Input[lineStart:position]) + 1
	caret := caretPadding(e.Input[lineStart:position]) + "^"

	return fmt.Sprintf("%s\n%s (%d:%d)\n%s", line, caret, lineNumber, column, e.Message)
}

func caretPadding(prefix string) string {
	var padding strings.Builder
	for _, r := range prefix {
		if r == '\t' {
			padding.WriteRune(r)
		} else {
			padding.WriteByte(' ')
		}
	}
	return padding.String()
}
