// generate README.md
package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/gnue/go2readme"
	"github.com/jessevdk/go-flags"
)

//go:generate go-bindata -ignore=\.DS_Store assets/...

var opts struct {
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

	err = d.WriteTo(os.Stdout)
	if err != nil {
		Fatalln(err)
	}
}

func ReadFileOrAsset(fname string, asset string) ([]byte, error) {
	if fname != "" {
		return ioutil.ReadFile(fname)
	}

	return Asset(asset)
}
