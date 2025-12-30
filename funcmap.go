package go2readme

import (
	"fmt"
	"os"
	"path/filepath"
)

type FuncMap map[string]interface{}

var DefualtFuncMap = FuncMap{
	"import": ImportFunc,
	"exists": ExistsFunc,
	"glob":   GlobFunc,
}

func ImportFunc(fname string) string {
	b, err := os.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return ""
	}

	return string(b)
}

func ExistsFunc(fname string) bool {
	_, err := os.Stat(fname)
	return err == nil
}

func GlobFunc(fname string) []string {
	files, _ := filepath.Glob(fname)
	return files
}
