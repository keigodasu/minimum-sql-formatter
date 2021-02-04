package minimum_sql_formatter

//select col1, col2 from table

var keywords = map[string]TokenType {
	"SELECT": SELECT,
	"FROM": FROM,
}

func LookupIdent(ident string) TokenType  {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENT = "IDENT"
	COMMA = ","
	SEMICOLON = ";"

	SELECT = "SELECT"
	FROM = "FROM"
)

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}