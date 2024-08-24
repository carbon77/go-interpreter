package main

import (
	"flag"
	"fmt"
	"go-interpreter/evaluator"
	"go-interpreter/lexer"
	"go-interpreter/object"
	"go-interpreter/parser"
	"go-interpreter/repl"
	"io"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	file := flag.String("f", "", "Enter filename")

	flag.Parse()

	if *file != "" {
		evalFile(*file, os.Stdout)
		return
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		currentUser.Username)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}

func evalFile(filename string, out io.Writer) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := string(file)
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	result := evaluator.Eval(program, env)

	out.Write([]byte(result.Inspect()))
}
