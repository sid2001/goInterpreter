package token 

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

type DataType struct {
	Size	int
	Type	TokenType
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers and literals
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	EQUALITY = "=="
	
	LT = "<"
	GT = ">"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	RSQUARE

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"

	KEYWORD = "KEYWORD"

	//DATATYPE
	DT_INT = "DT_INT"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

var dataTypes = map[string]DataType{
	"int"	: DataType{Size: 4, Type:"int"},
	"double": DataType{Size: 8, Type:"double"},
	"float"	: DataType{Size: 4, Type: "float"},
	"char"	: DataType{Size: 2, Type: "char"},
	"bool"	: DataType{Size: 1, Type:"bool"},
	"short" : DataType{Size: 2, Type:"short"},
}

var keywordList = [...]TokenType{
	"main","print","printf","auto", "double", "int", "const", "Struct", "short", "float", "unsigned", "break", "else", "long", "continue", "switch", "for", "signed", "void", "case", "enum", "register", "default", "typedef", "goto", "sizeof", "volatile", "char", "extern", "return", "do", "union", "if", "static", "while" }

func LookupKeyword(k string) bool {
	for i:=0 ; i<len(keywordList); i++ {
		if(k == string(keywordList[i])) {
			return true
		}
	}
	return false
}

func LookupDataType(k string) DataType {
	if dt, ok := dataTypes[k]; ok {
		return dt
	}
	return DataType{Size:0, Type:"None"}
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	if(LookupKeyword(ident)) {
		return KEYWORD
	}
	return IDENT
} 
