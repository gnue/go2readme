// generate README.md
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/gnue/go2readme"
	"github.com/jessevdk/go-flags"
)

//go:generate go-bindata -ignore=\.DS_Store assets/...

var opts struct {
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

	if opts.Write && opts.Output == "" {
		opts.Output = "README.md"
	}

	b, err := ReadFileOrAsset(opts.Template, "assets/README.md")
	if err != nil {
		Fatalln(err)
	}

	funcMap := template.FuncMap(go2readme.DefualtFuncMap)
	templ, err := template.New("README").Funcs(funcMap).Parse(string(b))
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

func ReadFileOrAsset(fname string, asset string) ([]byte, error) {
	if fname != "" {
		return ioutil.ReadFile(fname)
	}

	return Asset(asset)
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
