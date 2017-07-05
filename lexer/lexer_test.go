package lexer

import (
	"pupperscript/lexer"
	"pupperscript/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `floof five = 5;
floof ten = 10;

floof add = boop(x, y) {
  x + y;
};

!4 < 9 * 9 > 1/0;
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "floof"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "floof"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "floof"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "boop"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.INT, "4"},
		{token.LT, "<"},
		{token.INT, "9"},
		{token.ASTERISK, "*"},
		{token.INT, "9"},
		{token.GT, ">"},
		{token.INT, "1"},
		{token.SLASH, "/"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
