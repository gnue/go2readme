// generate from go code to README
package go2readme

import (
	"fmt"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/gnue/go2readme/mdfmt"
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
		files := append(ctx.TestGoFiles, ctx.XTestGoFiles...)

		for _, name := range files {
			f := filepath.Join(ctx.Dir, name)
			file, err := parser.ParseFile(fset, f, nil, parser.ParseComments)
			if err != nil {
				continue
			}

			examples = append(examples, doc.Examples(file)...)
		}

		d.cache.examples = NewExamples(d.fset, examples...)
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

func (d *Document) Description() string {
	return strings.TrimRightFunc(d.pkg.Doc, unicode.IsSpace)
}

func (d *Document) Usage() string {
	if d.IsCommand() {
		name := d.Name()

		p := os.Getenv("PATH")
		defer os.Setenv("PATH", p)

		os.Setenv("PATH", ".:"+p)

		_, err := exec.LookPath(name)
		if err != nil {
			return fmt.Sprintf("$ %s -h", name)
		}

		cmd := exec.Command(name, "-h")
		b, _ := cmd.CombinedOutput()
		usage := string(b)
		usage = strings.TrimPrefix(usage, "Usage:")
		usage = strings.TrimSpace(usage)
		return usage
	}

	return ""
}

func (d *Document) IsCommand() bool {
	return d.ctx.IsCommand()
}

func (d *Document) WriteTo(w io.Writer) error {
	mdw := mdfmt.NewWriter(w)

	err := d.templ.Execute(mdw, d)
	if err != nil {
		return err
	}

	return mdw.Flush()
}
