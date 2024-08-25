package evaluator

import (
	"fmt"
	"go-interpreter/object"
	"os"
)

var (
	builtins = map[string]*object.BuiltIn{
		"len": {
			Function: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1",
						len(args))
				}

				switch arg := args[0].(type) {
				case *object.String:
					return &object.Integer{Value: int64(len(arg.Value))}
				case *object.Array:
					return &object.Integer{Value: int64(len(arg.Elements))}
				default:
					return newError("argument to `len` not supported, got %s",
						args[0].Type())

				}
			},
		},
		"push": {
			Function: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, want=2",
						len(args))
				}

				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `push` must be ARRAY, got %s",
						args[0].Type())
				}

				arr := args[0].(*object.Array)
				length := len(arr.Elements)

				newElements := make([]object.Object, length+1, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &object.Array{Elements: newElements}
			},
		},
		"exit": {
			Function: func(args ...object.Object) object.Object {
				if len(args) != 0 {
					return newError("wrong number of arguments. got=%d, want=0", len(args))
				}

				os.Exit(0)
				return NULL
			},
		},
		"print": {
			Function: func(args ...object.Object) object.Object {
				for _, arg := range args {
					fmt.Println(arg.Inspect())
				}
				return NULL
			},
		},
	}
)
