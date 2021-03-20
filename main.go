package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil { panic(err) }

	fmt.Printf("Hello %s! these is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
