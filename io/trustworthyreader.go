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
	read, err = io.ReadFull(r.r, p)
	if err == io.ErrUnexpectedEOF {
		err = io.EOF
	}

	return
}
