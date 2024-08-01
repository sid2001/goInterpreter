package main

import (
//	"github.com/goInterpreter/repl"
	"os"
	"io/ioutil"
	"github.com/goInterpreter/token"
	"github.com/goInterpreter/lexer"
	//"fmt"
	"log"
)

func main() {
	
	file, err := os.Open("test")
	if err != nil {
        	log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

    // Convert byte slice to string
	content := string(data)
	lex := lexer.New(content)

	for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
		//fmt.Printf("%+v\n", tok)
		x := 2
		x++
	}

//	lex.PrintLexemes()
	lex.PrintSTable()
	//repl.Start(os.Stdin, os.Stdout)
}

