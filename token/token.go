package tokne

type Tokentype string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPAREN  = ")"
	LBREACE = "{"
	RBREACE = "}"

	//예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Tokentype{
	"fn":  FUNCTION,
	"let": LET,
}

func Lookupident(ident string) Tokentype {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
