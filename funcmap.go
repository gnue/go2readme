package go2readme

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FuncMap map[string]interface{}

var DefualtFuncMap = FuncMap{
	"import": ImportFunc,
}

func ImportFunc(fname string) string {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return ""
	}

	return string(b)
}
