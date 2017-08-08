package main

import (
	"fmt"
	"os"
)

func Fatal(v ...interface{}) {
	fmt.Fprint(os.Stderr, v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
