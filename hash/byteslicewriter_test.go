package hash

import (
	. "gopkg.in/check.v1"
)

type WrappedByteSliceSuite struct{}

var _ = Suite(&WrappedByteSliceSuite{})

func (s *WrappedByteSliceSuite) TestWrappedByteSlice(c *C) {
	buff := []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0xAA}
	w := NewSliceWriter(buff[1:])

	w.Write([]byte{0x01, 0x02})
	w.Write([]byte{0x04, 0x03})

	c.Assert(buff, DeepEquals, []byte{0xFF, 0x01, 0x02, 0x04, 0x03, 0xAA})
}
