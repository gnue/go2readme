package go2readme

import (
	"bytes"
	"go/doc"
	"go/format"
	"go/printer"
	"go/token"
	"strings"
)

type Example struct {
	fset *token.FileSet
	doc  *doc.Example
}

func NewExamples(fset *token.FileSet, examples ...*doc.Example) []*Example {
	list := make([]*Example, len(examples))

	for i, e := range examples {
		list[i] = &Example{fset: fset, doc: e}
	}

	return list
}

func (e *Example) Name() string {
	return e.doc.Name
}

func (e *Example) Doc() string {
	return e.doc.Doc
}

func (e *Example) Output() string {
	return e.doc.Output
}

func (e *Example) Code() string {
	var buf bytes.Buffer

	node := &printer.CommentedNode{Node: e.doc.Code, Comments: e.doc.Comments}

	config := &printer.Config{Mode: printer.UseSpaces, Tabwidth: 4}
	config.Fprint(&buf, e.fset, node)

	code := buf.String()
	if n := len(code); n >= 2 && code[0] == '{' && code[n-1] == '}' {
		code = code[1 : n-1]
		code = strings.Replace(code, "\n    ", "\n", -1)
		code = strings.Trim(code, " \r\n")
	}

	return code
}

func (e *Example) Play() string {
	if e.doc.Play == nil {
		return ""
	}

	var buf bytes.Buffer

	err := format.Node(&buf, e.fset, e.doc.Play)
	if err != nil {
		return ""
	}

	return buf.String()
}
