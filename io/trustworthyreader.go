package io

import (
	"io"
)

// TrustworthyReader implements "readExactly" semantics for an io.Reader object.
type trustworthyReader struct {
	r io.Reader
}

func NewTrustworthyReader(r io.Reader) io.Reader {
	return trustworthyReader{r}
}

// Reads exactly len(p) bytes at once, unless EOF is encountered.
func (r trustworthyReader) Read(p []byte) (read int, err error) {
	for read < len(p) && err == nil {
		var n int
		n, err = r.r.Read(p[read:])
		read += n
	}
	return
}
