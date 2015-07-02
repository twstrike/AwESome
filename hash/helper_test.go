package hash

import (
	"encoding/hex"
	"io"

	. "gopkg.in/check.v1"
)

type TestCase struct {
	input  string
	result string
	repeat int
}

type repeatReader struct {
	from  []byte
	at    int
	count int
}

func (st repeatReader) potentialEOF() error {
	if st.at == len(st.from) && st.count == 0 {
		return io.EOF
	}

	return nil
}

func (st *repeatReader) Read(into []byte) (n int, err error) {
	fromLen := len(st.from)
	numRead := copy(into, st.from[st.at:])
	st.at += numRead
	if st.at == fromLen && st.count > 0 {
		st.count--
		st.at = 0
	}
	return numRead, st.potentialEOF()
}

func repeat(b []byte, times int) *repeatReader {
	return &repeatReader{b, 0, times - 1}
}

func testOnSum(hash Hash, tc TestCase, c *C) {
	current := hash.Sum(repeat([]byte(tc.input), tc.repeat))
	r, _ := hex.DecodeString(tc.result)
	c.Assert(current, DeepEquals, r)
}
