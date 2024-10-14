package lexer

import (
	"fmt"
	"parser/constants"
	"unicode"
)

type Lexerer interface {
	NewToken() constants.Token
	readNumber() string
	readVariable() string
	peakChar() byte
	skipWhitespace()
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
}

func NewLexer(input string) Lexerer {
	l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

// readChar reads the next character in the input string and advances the position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) ToSring() string {
	return fmt.Sprintf("input: %s, position: %d, readPosition: %d, ch: %c, line: %d", l.input, l.position, l.readPosition, l.ch, l.line)
}

func (l *Lexer) NewToken() constants.Token {
	var tok constants.Token
	skipReadChar := false

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = constants.Token{Type: constants.TOKEN_DOUBLE_EQUAL, Lexeme: string(ch) + string(l.ch)}
		} else {
			tok = constants.Token{Type: constants.TOKEN_EQUAL, Lexeme: string(l.ch)}
		}
	case '<':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = constants.Token{Type: constants.TOKEN_NOT_EQUAL, Lexeme: string(ch) + string(l.ch)}
		}
	case '!':
		tok = constants.Token{Type: constants.TOKEN_NOT, Lexeme: string(l.ch)}
	case '+':
		tok = constants.Token{Type: constants.TOKEN_PLUS, Lexeme: string(l.ch)}
	case '-':
		tok = constants.Token{Type: constants.TOKEN_MINUS, Lexeme: string(l.ch)}
	case '*':
		tok = constants.Token{Type: constants.TOKEN_MULTIPLY, Lexeme: string(l.ch)}
	case '/':
		tok = constants.Token{Type: constants.TOKEN_DIVIDE, Lexeme: string(l.ch)}
	case '(':
		tok = constants.Token{Type: constants.TOKEN_LEFT_PAREN, Lexeme: string(l.ch)}
	case ')':
		tok = constants.Token{Type: constants.TOKEN_RIGHT_PAREN, Lexeme: string(l.ch)}
	case 0:
		tok.Lexeme = ""
		tok.Type = constants.TOKEN_EOF
		tok.Line = l.line
	default:
		if unicode.IsDigit(rune(l.ch)) {
			tok = constants.Token{
				Type:   constants.TOKEN_NUMBER,
				Lexeme: l.readNumber(),
			}
			skipReadChar = true
		} else if unicode.IsLetter(rune(l.ch)) {
			tok = constants.Token{
				Type:   constants.TOKEN_VARIABLE,
				Lexeme: l.readVariable(),
			}
			skipReadChar = true
		} else {
			tok = constants.Token{
				Type:   constants.TOKEN_EOF,
				Lexeme: string(l.ch),
			}
		}
	}

	tok.Line = l.line
	if !skipReadChar {
		l.readChar()
	}
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	seenDot := false
	for l.isDigit(l.ch) || (l.ch == '.' && !seenDot) {
		if l.ch == '.' {
			seenDot = true
		}
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readVariable() string {
	currentPosition := l.position

	for unicode.IsLetter(rune(l.ch)) {
		l.readChar()
	}

	return l.input[currentPosition:l.position]
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
