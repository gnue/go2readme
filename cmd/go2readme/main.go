// generate README for Go program
package main

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/gnue/go2readme"
	"github.com/jessevdk/go-flags"
)

//go:embed assets/README.md
var readme string

var opts struct {
	Version  bool   `short:"v" long:"version" description:"print version"`
	Write    bool   `short:"w" long:"write" description:"write to file"`
	Output   string `short:"o" long:"output" description:"output file"`
	Template string `short:"t" long:"template" description:"template file"`
	Args     struct {
		Dir string `positional-arg-name:"dir" default:"." description:"directory"`
	} `positional-args:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	if opts.Version {
		versionPrint()
		os.Exit(0)
	}

	if opts.Write && opts.Output == "" {
		opts.Output = "README.md"
	}

	funcMap := template.FuncMap(go2readme.DefualtFuncMap)
	templ, err := template.New("README").Funcs(funcMap).Parse(readme)
	if err != nil {
		Fatalln(err)
	}

	d, err := go2readme.NewDocument(opts.Args.Dir, templ)
	if err != nil {
		Fatalln(err)
	}

	w := os.Stdout

	if opts.Output != "" {
		f, err := createFile(opts.Output, false)
		if err != nil {
			Fatalln(err)
		}
		defer f.Close()

		w = f
	}

	err = d.WriteTo(w)
	if err != nil {
		Fatalln(err)
	}

	if opts.Output != "" {
		fmt.Fprintf(os.Stderr, "written to '%s'\n", opts.Output)
	}
}

func createFile(fname string, backup bool) (*os.File, error) {
	if backup {
		if _, err := os.Stat(fname); err == nil {
			err := os.Rename(fname, fname+".bak")
			if err != nil {
				return nil, err
			}
		}
	}

	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	if err := f.Truncate(0); err != nil {
		f.Close()
		return nil, err
	}

	return f, err
}
