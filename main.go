package main

import (
	"flag"
	"fmt"
	"go-interpreter/compiler"
	"go-interpreter/evaluator"
	"go-interpreter/lexer"
	"go-interpreter/object"
	"go-interpreter/parser"
	"go-interpreter/repl"
	"go-interpreter/vm"
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
	compilerMode := flag.Bool("c", false, "Use compiler")

	flag.Parse()

	if *file != "" {
		if *compilerMode {
			EvalFileWithCompiler(*file, os.Stdout)
			return
		}
		EvalFile(*file, os.Stdout)
		return
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		currentUser.Username)

	fmt.Printf("Feel free to type in commands\n")

	if *compilerMode {
		fmt.Printf("Compiler mode\n")
		repl.StartCompiler(os.Stdin, os.Stdout)
	} else {
		fmt.Printf("Interpreter mode\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}

func EvalFile(filename string, out io.Writer) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := string(file)
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	evaluated := evaluator.Eval(program, env)
	if err, ok := evaluated.(*object.Error); ok {
		fmt.Println(err.Inspect())
	}

}

func EvalFileWithCompiler(filename string, out io.Writer) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := string(file)
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	c := compiler.New()
	err = c.Compile(program)
	if err != nil {
		fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
		return
	}

	machine := vm.New(c.Bytecode())
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
		return
	}
	stackTop := machine.LastPoppedStackElem()
	io.WriteString(out, stackTop.Inspect())
	io.WriteString(out, "\n")
}
