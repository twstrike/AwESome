package hash

import "io"

// Hash is a hash function
type Hash interface {
	Sum(io.Reader) []byte
}
