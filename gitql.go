package main

import (
	"github.com/cloudson/gitql/parser"
	"github.com/cloudson/gitql/runtime"
	"github.com/cloudson/gitql/semantical"
	"github.com/nemith/goline"
	"log"
	"path/filepath"
)

// global vars set in init() function 
var query string
var interactive *bool

func main() {
	folder, errFile := filepath.Abs(*path)

	if errFile != nil {
		log.Fatalln(errFile)
	}

	if !*interactive {
		runQuery(query, folder)
		return 
	}

	gl := goline.NewGoLine(goline.StringPrompt("gitql> "))
	for {
		data, err := gl.Line()
		if err != nil {
        	if err != goline.UserTerminatedError {
        		log.Fatalln(err)
        	}
        }

		if err == goline.UserTerminatedError || data == "exit" || data == "quit" {
            return
        }
        runQuery(data, folder)
	}
}

func runQuery(query, folder string) {
	parser.New(query)
	ast, errGit := parser.AST()
	if errGit != nil {
		log.Fatalln(errGit)
	}
	ast.Path = &folder
	errGit = semantical.Analysis(ast)
	if errGit != nil {
		log.Fatalln(errGit)
	}

	runtime.Run(ast)
}
