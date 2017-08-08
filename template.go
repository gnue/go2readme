package go2readme

import (
	"io"
)

type Template interface {
	Execute(w io.Writer, data interface{}) error
}
