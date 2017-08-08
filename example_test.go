package go2readme

import (
	"bytes"
	"fmt"
	"text/template"
)

// Document output
func ExampleDocument() {
	// read template
	templ := template.Must(template.ParseFiles("template.md"))
	d, err := NewDocument(".", templ)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	err = d.WriteTo(&buf)
	if err != nil {
		return
	}

	fmt.Println(buf.String())
}
