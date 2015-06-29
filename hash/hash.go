package hash

import "io"

type Hash interface {
	Sum(io.Reader) []byte
}
