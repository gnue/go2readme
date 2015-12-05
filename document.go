// generate from go code to README
package go2readme

import (
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Document struct {
	ctx   *build.Package
	pkg   *doc.Package
	fset  *token.FileSet
	templ Template

	cache struct {
		examples []*Example
	}
}

func NewDocument(dir string, templ Template) (*Document, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	ctx, err := build.ImportDir(path, build.ImportComment)
	if err != nil {
		return nil, err
	}

	d := &Document{fset: token.NewFileSet(), ctx: ctx, templ: templ}
	err = d.parse()

	return d, err
}

func (d *Document) parse() error {
	ctx := d.ctx

	filter := func(file os.FileInfo) bool {
		return !strings.HasSuffix(file.Name(), "_test.go")
	}

	pkgs, err := parser.ParseDir(d.fset, ctx.Dir, filter, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, p := range pkgs {
		d.pkg = doc.New(p, ctx.ImportPath, 0)
		break
	}

	return nil
}

func (d *Document) Examples() []*Example {
	if d.cache.examples == nil {
		var examples []*doc.Example

		ctx := d.ctx
		fset := d.fset

		for _, name := range ctx.TestGoFiles {
			f := filepath.Join(ctx.Dir, name)
			file, err := parser.ParseFile(fset, f, nil, parser.ParseComments)
			if err != nil {
				continue
			}

			examples = append(examples, doc.Examples(file)...)
		}

		d.cache.examples = make([]*Example, len(examples))

		for i, e := range examples {
			d.cache.examples[i] = &Example{fset: d.fset, doc: e}
		}
	}

	return d.cache.examples
}

func (d *Document) Name() string {
	if d.IsCommand() {
		return filepath.Base(d.ctx.Dir)
	}

	return d.pkg.Name
}

func (d *Document) ImportPath() string {
	return d.pkg.ImportPath
}

func (d *Document) Synopsis() string {
	return doc.Synopsis(d.pkg.Doc)
}

func (d *Document) Usage() string {
	return ""
}

func (d *Document) IsCommand() bool {
	return d.ctx.IsCommand()
}

func (d *Document) WriteTo(w io.Writer) error {
	return d.templ.Execute(w, d)
}
