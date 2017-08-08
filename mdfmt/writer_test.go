package mdfmt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	sample := "\n# Title \n\ndocument... \n\n\ncontens\n...\nabc"
	result := "\n# Title\n\ndocument...\n\ncontens\n...\nabc"
	var b bytes.Buffer

	w := NewWriter(&b)
	_, err := w.Write([]byte(sample))
	assert.Nil(t, err)
	err = w.Flush()
	assert.Nil(t, err)
	assert.Equal(t, b.String(), result)
}

func TestWriteInTab(t *testing.T) {
	sample := "\n# Title \n\n\tdocument... \t\n\t\n\t\n\tcontens\n...\nabc"
	result := "\n# Title\n\n\tdocument...\n\n\n\tcontens\n...\nabc"
	var b bytes.Buffer

	w := NewWriter(&b)
	_, err := w.Write([]byte(sample))
	assert.Nil(t, err)
	err = w.Flush()
	assert.Nil(t, err)
	assert.Equal(t, b.String(), result)
}

func TestWriteInIndent(t *testing.T) {
	sample := "\n# Title \n\n    document...     \n    \n    \n    contens\n...\nabc"
	result := "\n# Title\n\n    document...\n\n\n    contens\n...\nabc"
	var b bytes.Buffer

	w := NewWriter(&b)
	_, err := w.Write([]byte(sample))
	assert.Nil(t, err)
	err = w.Flush()
	assert.Nil(t, err)
	assert.Equal(t, b.String(), result)
}

func TestWriteInBlockQuote(t *testing.T) {
	sample := "\n# Title \n\n ```document... \n\n\ncontens\n ```...\nabc"
	result := "\n# Title\n\n ```document...\n\n\ncontens\n ```...\nabc"
	var b bytes.Buffer

	w := NewWriter(&b)
	_, err := w.Write([]byte(sample))
	assert.Nil(t, err)
	err = w.Flush()
	assert.Nil(t, err)
	assert.Equal(t, b.String(), result)
}

func TestWriteInIndentNG(t *testing.T) {
	sample := "\n# Title \n    document...     \n    \n    \n    contens\n...\nabc"
	result := "\n# Title\n    document...\n\n    contens\n...\nabc"
	var b bytes.Buffer

	w := NewWriter(&b)
	_, err := w.Write([]byte(sample))
	assert.Nil(t, err)
	err = w.Flush()
	assert.Nil(t, err)
	assert.Equal(t, b.String(), result)
}
