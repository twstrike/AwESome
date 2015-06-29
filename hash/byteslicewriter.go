package hash

import "io"

type wrappedByteSlice struct {
	slice []byte
	from  int
}

func NewSliceWritter(b []byte) io.Writer {
	return &wrappedByteSlice{slice: b}
}

func (bs *wrappedByteSlice) Write(p []byte) (n int, err error) {
	copy(bs.slice[bs.from:], p)
	bs.from += len(p)
	return len(p), nil
}
