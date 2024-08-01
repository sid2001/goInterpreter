package lexer

import (
	"github.com/goInterpreter/token"
	"github.com/goInterpreter/sTable"
	"fmt"
	"strconv"
)



//type variableList []struct{
//	name		string
//	scope		string
//	blockName	string
//}

var AddrPointer = 1001

type Lexemes struct {
	variables	[]string
	literals	[]string
	operators	[]string
	constants	[]string
	keywords	[]string
	specialSymbols	[]string
}

type Lexer struct {
	input		string
	position	int // current position in input
	readPosition	int // current reading position in input
	ch		byte // current char under examination
	sTable		sTable.STable
	lexemes		Lexemes
	bufferDataType	token.DataType
}



func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) PrintSTable() {
	fmt.Println("STable:")
	for _, data := range l.sTable {
		fmt.Printf("Name: %s, Type: %s, Size: %d, Addr: %d\n", data.Name, data.Type, data.Size, data.Addr)
    }
}

func (l *Lexer) PrintLexemes() {
    fmt.Println("Variables:", l.lexemes.variables)
    fmt.Println("Literals:", l.lexemes.literals)
    fmt.Println("Operators:", l.lexemes.operators)
    fmt.Println("Constants:", l.lexemes.constants)
    fmt.Println("Keywords:", l.lexemes.keywords)
    fmt.Println("Special Symbols:", l.lexemes.specialSymbols)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for null
	} else {
		l.ch = l.input[l.readPosition]
	}
	
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	//fmt.Println(l.ch)
	switch l.ch {
	case '-':
		tok = newToken(token.MINUS, l.ch)
		l.lexemes.operators = append(l.lexemes.operators, string(l.ch))
	case '!':
		tok = newToken(token.BANG, l.ch)
		l.lexemes.operators = append(l.lexemes.operators, string(l.ch))
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
		l.lexemes.operators =append(l.lexemes.operators, string(l.ch))

	case '/':
		tok = newToken(token.SLASH, l.ch)
		l.lexemes.operators =append(l.lexemes.operators, string(l.ch))
	case '<':
		tok = newToken(token.LT, l.ch)
		l.lexemes.operators =append(l.lexemes.operators, string(l.ch))
	case '>':
		tok = newToken(token.GT, l.ch)
		l.lexemes.operators =append(l.lexemes.operators, string(l.ch))

	case '=':
		//fmt.Print("equal")
		if(l.peekForward(1) == "=") {
			tok.Type = token.EQUALITY
			tok.Literal = "=="
			l.seekForward(2)
		} else { 
			tok = newToken(token.ASSIGN, l.ch)
		}
		l.lexemes.operators =append(l.lexemes.operators, string(l.ch))

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		l.lexemes.specialSymbols =append(l.lexemes.specialSymbols, string(l.ch))

	case '(':
		tok = newToken(token.LPAREN, l.ch)
		l.lexemes.specialSymbols = append(l.lexemes.specialSymbols, string(l.ch))

	case ')':
		tok = newToken(token.RPAREN, l.ch)
		l.lexemes.specialSymbols =append(l.lexemes.specialSymbols, string(l.ch))

	case ',':
		tok = newToken(token.COMMA, l.ch)
		l.lexemes.specialSymbols = append(l.lexemes.specialSymbols, string(l.ch))

	case '+':
		tok = newToken(token.PLUS, l.ch)
		l.lexemes.operators = append(l.lexemes.operators, string(l.ch))

	case '{':
		tok = newToken(token.LBRACE, l.ch)
		l.lexemes.specialSymbols = append(l.lexemes.specialSymbols, string(l.ch))

	case '}':
		tok = newToken(token.RBRACE, l.ch)
		l.lexemes.specialSymbols = append(l.lexemes.specialSymbols, string(l.ch))

	case '"':
		tok.Literal = l.seekLiteral()
		tok.Type = "LITERAL"
		l.lexemes.literals = append(l.lexemes.literals, tok.Literal)
		return tok
	case 0:
		//fmt.Println("eof")
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if(isDigit(l.ch)){
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			l.lexemes.constants = append(l.lexemes.constants,tok.Literal)
			return tok
		} else if(isLetter(l.ch)) {
			//fmt.Println(l.ch)
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			
			if(l.isFunction()) {
				l.lexemes.variables = append(l.lexemes.variables,tok.Literal)
				dd := sTable.STableData{Name:string(tok.Literal),Type:"function",Size:0,Addr:0}
				l.sTable = append(l.sTable,dd)
				tt := token.DataType{Size:0,Type:token.TokenType("function")}
				varAttr := sTable.VariableAttributes{Addr:0,Type:tt}
				sTable.Variables[tok.Literal] = varAttr
			}else if(tok.Type == token.KEYWORD){
				l.lexemes.keywords = append(l.lexemes.keywords,tok.Literal)
				if dt := token.LookupDataType(tok.Literal) ; dt.Type != "None"{
					l.bufferDataType = dt
				}
			} else {
				if _, ok:= sTable.Variables[tok.Literal]; !ok{
					
					var arrSize int
					arrSize = l.checkArray()
					
					ty := string(l.bufferDataType.Type)
					siz := l.bufferDataType.Size

					if(arrSize != -1) {
						ty += "[]"
						siz = siz * arrSize
					}
					dd := sTable.STableData{Name: string(tok.Literal), Type:ty, Size:siz,Addr: AddrPointer}
					l.sTable = append(l.sTable, dd)
					l.lexemes.variables = append(l.lexemes.variables,tok.Literal)
					varAttr := sTable.VariableAttributes{
						Addr: AddrPointer,
						Type: l.bufferDataType}
					sTable.Variables[tok.Literal] = varAttr
					AddrPointer += l.bufferDataType.Size
					l.bufferDataType = token.DataType{Size:0,Type:"None"}

				}
			}
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	//fmt.Println(l.input[position:l.position])
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) seekForward(r int) {
	for r > 0 {
		r--
		l.readChar()
	}
	return
} 

func (l *Lexer) checkArray() int {
	var si = ""
	if(l.ch == '['){
		l.readChar()
		for(l.ch != ']') {
			si += string(l.ch)
			l.readChar()
		}
		l.readChar()
		x, _ := strconv.Atoi(si)
		return int(x);
	}
	return -1
}

func (l *Lexer) peekForward(r int) (str string) {
	str = ""

	if(l.position + r < len(l.input)) {
		str = l.input[l.position +1: l.position + 1 + r]
	}

	return 
}

func (l *Lexer) isFunction() bool {
	if(l.ch == '(') {
		return true
	}
	return false
}

func (l *Lexer) seekLiteral() string {
	position := l.position
	l.readChar()
	for l.ch != '"' {
		l.readChar()
	}
	l.readChar()
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
