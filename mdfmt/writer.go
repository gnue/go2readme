package mdfmt

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type state int

const (
	stateStart state = iota
	stateNormal
	stateBlankLine
	stateIndent
	stateBlockQuote
)

type Writer struct {
	out    io.Writer
	buf    []byte
	state  state
	blockq string
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{out: w}
}

func (w *Writer) Write(b []byte) (n int, err error) {
	buf := append(w.buf, b...)
	out := w.out

	defer func() {
		w.buf = buf
	}()

	for {
		advance, line, err := bufio.ScanLines(buf, false)
		buf = buf[advance:]
		if err != nil {
			return n, err
		}
		if advance == 0 {
			break
		}

		s := strings.TrimRightFunc(string(line), unicode.IsSpace)
		if w.state == stateBlankLine && s == "" {
			continue
		}

		switch w.state {
		case stateIndent:
			if !strings.HasPrefix(s, "\t") {
				w.state = stateNormal
			}
		case stateBlockQuote:
			if s == w.blockq {
				w.state = stateNormal
			}
		default:
			w.state = w.nextState(s)
		}

		size, err := out.Write([]byte(s + "\n"))
		n += size
		if err != nil {
			return n, err
		}
	}

	return
}

func (w *Writer) Flush() error {
	defer func() {
		w.buf = nil
	}()

	_, err := w.out.Write(w.buf)
	return err
}

func (w *Writer) nextState(s string) state {
	switch {
	case s == "":
		return stateBlankLine
	}

	switch w.state {
	case stateStart, stateBlankLine:
		if strings.HasPrefix(s, "\t") || strings.HasPrefix(s, "    ") {
			return stateIndent
		}
		if bq, ok := hasBlockQuote(s); ok {
			w.blockq = bq
			return stateBlockQuote
		}
	}

	return stateNormal
}

func hasBlockQuote(s string) (string, bool) {
	indent := strings.IndexFunc(s, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
	if indent < 0 {
		return "", false
	}

	const bq = "```"
	if strings.HasPrefix(s[indent:], bq) {
		return s[:indent] + bq, true
	}

	return "", false
}
