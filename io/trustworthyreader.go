package io

import (
	"io"
)

// TrustworthyReader implements "readExactly" semantics for an io.Reader object.
type TrustworthyReader struct {
	r io.Reader
}

// NewTrustworthyReader creates a new TrustworthyReader
func NewTrustworthyReader(r io.Reader) *TrustworthyReader {
	return &TrustworthyReader{r}
}

// Reads exactly len(p) bytes at once, unless EOF is encountered.
func (r TrustworthyReader) Read(p []byte) (read int, err error) {
	for read < len(p) && err == nil {
		var n int
		n, err = r.r.Read(p[read:])
		read += n
	}
	return
}
