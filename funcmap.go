package go2readme

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FuncMap map[string]interface{}

var DefualtFuncMap = FuncMap{
	"import": ImportFunc,
	"exists": ExistsFunc,
}

func ImportFunc(fname string) string {
	b, err := ioutil.ReadFile(fname)
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
