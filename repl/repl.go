package repl

import (
	"bufio"
	"io"
	"fmt"
	"github.com/goInterpreter/lexer"
	"github.com/goInterpreter/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		text := scanner.Text()
		
		lex := lexer.New(text)
		
		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
	
}
