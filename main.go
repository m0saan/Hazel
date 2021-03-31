package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {

	usr, err := user.Current()
	if err != nil { panic(err) }

	fmt.Printf("Hello %s! these is the Monkey programming language!\n", usr.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
