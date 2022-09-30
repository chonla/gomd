package main

import (
	"fmt"
	"gomd/parser"
	"gomd/reader"

	"github.com/kr/pretty"
)

func main() {
	rd := reader.New()
	lines, e := rd.ReadTextFile("./examples/md/example-2.md")
	if e != nil {
		fmt.Println(e)
	}

	doc, e := parser.Parse(lines)
	if e != nil {
		fmt.Println(e)
	}

	pretty.Println(doc)
}
