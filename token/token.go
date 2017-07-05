package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	ASTERISK  = "*"
	SLASH     = "/"
	BANG      = "!"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LT        = "<"
	GT        = ">"
	LTE       = "<="
	GTE       = ">="
	LET       = "LET"
	FUNCTION  = "FUNCTION"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
)

var keywords = map[string]TokenType{
	"boop":  FUNCTION,
	"floof": LET,
	"such":  IF,
	"so":    ELSE,
	"bork":  RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
