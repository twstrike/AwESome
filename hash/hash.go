package hash

type Hash interface {
	Init()
	Sum([]byte) []byte
}
