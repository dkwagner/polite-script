package main

import (
	"fmt"
	"os"
	"pscript/repl"
)

const pscriptLogo = ` ___     _ _ _         ___        _      _    
| _ \___| (_| |_ ___  / __|__ _ _(_)_ __| |_  
|  _/ _ | | |  _/ -_) \__ / _| '_| | '_ |  _| 
|_| \___|_|_|\__\___| |___\__|_| |_| .__/\__| 
				   |_|        `

func main() {
	fmt.Printf("%s\n\nHello, and welcome to polite-script!\n(To exit type 'exit' or ctrl+c)\n", pscriptLogo)

	repl.Start(os.Stdin, os.Stdout)
}
