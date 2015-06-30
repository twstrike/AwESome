package hash

import (
	"bytes"
	"io"

	. "gopkg.in/check.v1"
)

func (s *SHA1Suite) TestReadExactly(c *C) {
	input := []byte{0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12}
	reader := bytes.NewBuffer(input)
	into := [5]byte{}

	n, err := readExactly(into[:], reader)
	c.Assert(into[:], DeepEquals, input[:5])
	c.Assert(n, Equals, len(into))
	c.Assert(err, DeepEquals, nil)
}

func (s *SHA1Suite) TestReadExactlyWithEOFError(c *C) {
	input := []byte{0x11, 0x12, 0x11, 0x12}
	reader := bytes.NewBuffer(input)
	into := [5]byte{}

	n, err := readExactly(into[:], reader)
	c.Assert(into[:], DeepEquals, []byte{0x11, 0x12, 0x11, 0x12, 0x00})
	c.Assert(n, Equals, len(input))
	c.Assert(err, DeepEquals, io.EOF)
}
