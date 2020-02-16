package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dkwagner/pscript/lexer"
	"github.com/dkwagner/pscript/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		// Allow user to exit
		if line == "exit" {
			os.Exit(0)
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
