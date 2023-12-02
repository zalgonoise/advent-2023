package parser

import "github.com/zalgonoise/parse"

func ParseFunc(t *parse.Tree[Token, byte]) parse.ParseFn[Token, byte] {
	switch t.Peek().Type {
	case TokenAlpha:
		return parseAlpha
	case TokenColon:
		return parseColon
	case TokenNum:
		return parseNum
	case TokenComma:
		return parseComma
	case TokenSemicolon:
		return parseSemicolon
	case TokenNewline, TokenEOF, TokenError:
		return nil
	default:
		return nil
	}
}

func parseAlpha(t *parse.Tree[Token, byte]) parse.ParseFn[Token, byte] {
	parent := t.Node(t.Next())

	if t.Peek().Type == TokenSpace {
		_ = t.Next()
	}

	if t.Peek().Type != TokenNum {
		return nil
	}

	t.Node(t.Next())
	_ = t.Set(parent)

	return ParseFunc
}

func parseColon(t *parse.Tree[Token, byte]) parse.ParseFn[Token, byte] {
	t.Node(t.Next())

	if t.Peek().Type == TokenSpace {
		_ = t.Next()
	}

	return ParseFunc
}

func parseNum(t *parse.Tree[Token, byte]) parse.ParseFn[Token, byte] {
	//parent := t.Node(t.Next())
	t.Node(t.Next())

	if t.Peek().Type == TokenSpace {
		_ = t.Next()
	}

	if t.Peek().Type == TokenAlpha {
		t.Node(t.Next())
	}

	//_ = t.Set(parent)

	return ParseFunc
}

func parseComma(t *parse.Tree[Token, byte]) parse.ParseFn[Token, byte] {
	parent := t.Cur()

	for parent.Parent.Type != TokenColon && parent.Parent.Type != TokenSemicolon {
		if parent.Parent == nil {
			return nil
		}

		parent = parent.Parent
	}

	_ = t.Set(parent)
	t.Node(t.Next())

	if t.Peek().Type == TokenSpace {
		_ = t.Next()
	}

	if t.Peek().Type != TokenNum {
		return nil
	}

	return parseNum
}

func parseSemicolon(t *parse.Tree[Token, byte]) parse.ParseFn[Token, byte] {
	parent := t.Cur()

	for parent.Type != TokenColon {
		if parent.Parent == nil {
			return nil
		}

		parent = parent.Parent
	}

	_ = t.Set(parent)
	t.Node(t.Next())

	if t.Peek().Type == TokenSpace {
		t.Next()
	}

	if t.Peek().Type != TokenNum {
		return nil
	}

	return parseNum
}
