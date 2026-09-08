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
	"errors"
	"reflect"
	"testing"
)

func token(tokenType TokenType, value string, start, end int) Token {
	return Token{Type: tokenType, Value: value, Start: start, End: end}
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Token
		wantErr bool
	}{
		{
			name:  "empty input",
			input: "",
			want:  []Token{},
		},
		{
			name:  "whitespace",
			input: " \t\n",
			want:  []Token{},
		},
		{
			name:  "single identifier",
			input: "priority",
			want:  []Token{token(IDENTIFIER, "priority", 0, 8)},
		},
		{
			name:  "multiple identifiers",
			input: "priority medium high",
			want: []Token{
				token(IDENTIFIER, "priority", 0, 8),
				token(IDENTIFIER, "medium", 9, 15),
				token(IDENTIFIER, "high", 16, 20),
			},
		},
		{
			name:  "newline separated identifiers",
			input: "priority\nmedium",
			want: []Token{
				token(IDENTIFIER, "priority", 0, 8),
				token(IDENTIFIER, "medium", 9, 15),
			},
		},
		{
			name:  "string",
			input: `"brother"`,
			want:  []Token{token(STRING, `"brother"`, 0, 9)},
		},
		{
			name:  "string with spaces",
			input: `"call my brother"`,
			want:  []Token{token(STRING, `"call my brother"`, 0, 17)},
		},
		{
			name:  "empty string",
			input: `""`,
			want:  []Token{token(STRING, `""`, 0, 2)},
		},
		{
			name:  "string with escaped quote",
			input: `"say \"hello\""`,
			want:  []Token{token(STRING, `"say \"hello\""`, 0, 15)},
		},
		{
			name:  "and",
			input: "and",
			want:  []Token{token(AND, "and", 0, 3)},
		},
		{
			name:  "or",
			input: "or",
			want:  []Token{token(OR, "or", 0, 2)},
		},
		{
			name:  "not",
			input: "not",
			want:  []Token{token(NOT, "not", 0, 3)},
		},
		{
			name:  "not operator",
			input: "!",
			want:  []Token{token(NOT, "!", 0, 1)},
		},
		{
			name:  "case insensitive keywords preserve source spelling",
			input: "AND Or nOt HaS",
			want: []Token{
				token(AND, "AND", 0, 3),
				token(OR, "Or", 4, 6),
				token(NOT, "nOt", 7, 10),
				token(HAS, "HaS", 11, 14),
			},
		},
		{
			name:  "has",
			input: "has",
			want:  []Token{token(HAS, "has", 0, 3)},
		},
		{
			name:  "equals",
			input: "=",
			want:  []Token{token(EQUALS, "=", 0, 1)},
		},
		{
			name:  "not equals",
			input: "!=",
			want:  []Token{token(NOT_EQUALS, "!=", 0, 2)},
		},
		{
			name:  "greater than",
			input: ">",
			want:  []Token{token(GREATER_THAN, ">", 0, 1)},
		},
		{
			name:  "less than",
			input: "<",
			want:  []Token{token(LESS_THAN, "<", 0, 1)},
		},
		{
			name:  "greater or equal",
			input: ">=",
			want:  []Token{token(GREATER_OR_EQUAL, ">=", 0, 2)},
		},
		{
			name:  "less or equal",
			input: "<=",
			want:  []Token{token(LESS_OR_EQUAL, "<=", 0, 2)},
		},
		{
			name:  "fuzzy match",
			input: "~=",
			want:  []Token{token(FUZZY_MATCH, "~=", 0, 2)},
		},
		{
			name:  "parentheses",
			input: "()",
			want: []Token{
				token(LEFT_PAREN, "(", 0, 1),
				token(RIGHT_PAREN, ")", 1, 2),
			},
		},
		{
			name:  "complete comparison",
			input: "priority > medium",
			want: []Token{
				token(IDENTIFIER, "priority", 0, 8),
				token(GREATER_THAN, ">", 9, 10),
				token(IDENTIFIER, "medium", 11, 17),
			},
		},
		{
			name:  "complete query",
			input: `priority > medium and title has "brother"`,
			want: []Token{
				token(IDENTIFIER, "priority", 0, 8),
				token(GREATER_THAN, ">", 9, 10),
				token(IDENTIFIER, "medium", 11, 17),
				token(AND, "and", 18, 21),
				token(IDENTIFIER, "title", 22, 27),
				token(HAS, "has", 28, 31),
				token(STRING, `"brother"`, 32, 41),
			},
		},
		{
			name:  "query with parentheses",
			input: `(priority > medium)`,
			want: []Token{
				token(LEFT_PAREN, "(", 0, 1),
				token(IDENTIFIER, "priority", 1, 9),
				token(GREATER_THAN, ">", 10, 11),
				token(IDENTIFIER, "medium", 12, 18),
				token(RIGHT_PAREN, ")", 18, 19),
			},
		},
		{
			name:  "operators without whitespace",
			input: `priority>=medium`,
			want: []Token{
				token(IDENTIFIER, "priority", 0, 8),
				token(GREATER_OR_EQUAL, ">=", 8, 10),
				token(IDENTIFIER, "medium", 10, 16),
			},
		},
		{
			name:  "all comparison operators",
			input: `= != > >= < <=`,
			want: []Token{
				token(EQUALS, "=", 0, 1),
				token(NOT_EQUALS, "!=", 2, 4),
				token(GREATER_THAN, ">", 5, 6),
				token(GREATER_OR_EQUAL, ">=", 7, 9),
				token(LESS_THAN, "<", 10, 11),
				token(LESS_OR_EQUAL, "<=", 12, 14),
			},
		},
		{
			name:  "unicode identifiers use byte offsets",
			input: "проект>=высокий",
			want: []Token{
				token(IDENTIFIER, "проект", 0, 12),
				token(GREATER_OR_EQUAL, ">=", 12, 14),
				token(IDENTIFIER, "высокий", 14, 28),
			},
		},
		{
			name:  "replacement character is a valid identifier",
			input: "�",
			want:  []Token{token(IDENTIFIER, "�", 0, 3)},
		},
		{
			name:    "unterminated string",
			input:   `"brother`,
			wantErr: true,
		},
		{
			name:    "unterminated escaped string",
			input:   "\"brother\\",
			wantErr: true,
		},
		{
			name:    "invalid UTF-8",
			input:   string([]byte{0xff}),
			wantErr: true,
		},
		{
			name:    "unpaired fuzzy match operator",
			input:   "~",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("Tokenize() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Tokenize() unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() mismatch\n got:  %#v\n want: %#v", got, tt.want)
			}
		})
	}
}

func TestTokenizeErr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		position int
		message  string
		want     string
	}{
		{
			name:     "unterminated string on first line",
			input:    `"brother`,
			position: 0,
			message:  "unterminated string literal",
			want:     "\"brother\n^ (1:1)\nunterminated string literal",
		},
		{
			name:     "unterminated string on later line",
			input:    "priority\n  \"brother",
			position: 11,
			message:  "unterminated string literal",
			want:     "  \"brother\n  ^ (2:3)\nunterminated string literal",
		},
		{
			name:     "invalid UTF-8",
			input:    string([]byte{0xff}),
			position: 0,
			message:  "invalid UTF-8",
			want:     string([]byte{0xff}) + "\n^ (1:1)\ninvalid UTF-8",
		},
		{
			name:     "unpaired fuzzy match operator",
			input:    "priority ~ high",
			position: 9,
			message:  "expected '=' after '~'",
			want:     "priority ~ high\n         ^ (1:10)\nexpected '=' after '~'",
		},
		{
			name:     "caret preserves tabs",
			input:    "priority\t~",
			position: 9,
			message:  "expected '=' after '~'",
			want:     "priority\t~\n        \t^ (1:10)\nexpected '=' after '~'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Tokenize(tt.input)
			if err == nil {
				t.Fatal("Tokenize() expected an error, got nil")
			}

			var tokenizeErr TokenizeErr
			if !errors.As(err, &tokenizeErr) {
				t.Fatalf("Tokenize() error type = %T, want TokenizeErr", err)
			}
			if tokenizeErr.Input != tt.input || tokenizeErr.Position != tt.position || tokenizeErr.Message != tt.message {
				t.Errorf("TokenizeErr = %#v, want Input: %q, Position: %d, Message: %q", tokenizeErr, tt.input, tt.position, tt.message)
			}
			if got := tokenizeErr.Error(); got != tt.want {
				t.Errorf("TokenizeErr.Error() = %q, want %q", got, tt.want)
			}
		})
	}
}
