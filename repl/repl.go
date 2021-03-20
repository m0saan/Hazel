package repl

import (
	"Interpreter/lexer"
	"Interpreter/token"
	"bufio"
	"fmt"
	"io"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for true {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned{ return }
		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
			fmt.Printf("%+v\n", tok)
		}
	}

}