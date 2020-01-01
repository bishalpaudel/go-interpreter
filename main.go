package main

import (
	"fmt"
	// "os"
	// "os/user"
	// "go-interpreter/repl"
	"go-interpreter/ast"
	"go-interpreter/token"
)

func main() {
	Statements := []ast.Statement{}

	t := token.Token{token.EOF, ""}
	i := &ast.Identifier{t, "justRandom"}

	statement1 := &ast.LetStatement{Token: t}

	append(Statements, statement1)

	fmt.Printf(Statements)




	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Hello %s! This is the Monkey programming language.\n", user.Username)
	// fmt.Printf("Feel free to type in commands\n")
	// repl.Start(os.Stdin, os.Stdout)
}