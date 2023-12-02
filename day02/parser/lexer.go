package parser

import (
	"github.com/zalgonoise/lex"
)

type Token uint8

const (
	TokenEOF Token = iota
	TokenError
	TokenAlpha
	TokenNum
	TokenColon
	TokenSemicolon
	TokenComma
	TokenNewline
	TokenSpace
)

func StateFunc(l lex.Lexer[Token, byte]) lex.StateFn[Token, byte] {
	switch item := l.Next(); {
	case item == ':':
		l.Emit(TokenColon)

		return StateFunc
	case item == ';':
		l.Emit(TokenSemicolon)

		return StateFunc
	case item == ',':
		l.Emit(TokenComma)

		return StateFunc
	case item == '\n':
		l.Emit(TokenNewline)

		return StateFunc
	case item == ' ':
		l.Emit(TokenSpace)

		return StateFunc
	case item == 0:
		l.Emit(TokenEOF)

		return nil
	case item >= '0' && item <= '9':
		return stateNum
	case (item >= 'A' && item <= 'Z') || (item >= 'a' && item <= 'z'):
		return stateAlpha
	default:
		l.Emit(TokenError)

		return nil
	}
}

func stateAlpha(l lex.Lexer[Token, byte]) lex.StateFn[Token, byte] {
	l.Backup() // undo l.Next() for the l.AcceptRun call

	for {
		if item := l.Cur(); (item >= 'A' && item <= 'Z') || (item >= 'a' && item <= 'z') {
			l.Next()

			continue
		}

		break
	}

	if l.Width() > 0 {
		l.Emit(TokenAlpha)
	}

	return StateFunc
}

func stateNum(l lex.Lexer[Token, byte]) lex.StateFn[Token, byte] {
	l.Backup() // undo l.Next() for the l.AcceptRun call

	for {
		if item := l.Cur(); item >= '0' && item <= '9' {
			l.Next()

			continue
		}

		break
	}

	if l.Width() > 0 {
		l.Emit(TokenNum)
	}

	return StateFunc
}
