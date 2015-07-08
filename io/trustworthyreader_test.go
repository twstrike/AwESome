package io

import (
	"bytes"
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TrustworthyReaderSuite struct{}

var _ = Suite(&TrustworthyReaderSuite{})

func (s *TrustworthyReaderSuite) TestReadExactly(c *C) {
	input := []byte{0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12, 0x11, 0x12}
	reader := NewTrustworthyReader(bytes.NewBuffer(input))
	into := [5]byte{}

	n, err := reader.Read(into[:])
	c.Assert(into[:], DeepEquals, input[:5])
	c.Assert(n, Equals, len(into))
	c.Assert(err, DeepEquals, nil)
}

func (s *TrustworthyReaderSuite) TestReadExactlyWithEOFError(c *C) {
	input := []byte{0x11, 0x12, 0x11, 0x12}
	reader := NewTrustworthyReader(bytes.NewBuffer(input))
	into := [5]byte{}

	n, err := reader.Read(into[:])
	c.Assert(into[:], DeepEquals, []byte{0x11, 0x12, 0x11, 0x12, 0x00})
	c.Assert(n, Equals, len(input))
	c.Assert(err, DeepEquals, io.EOF)
}

func (s *TrustworthyReaderSuite) TestReadExactlySubsequentCalls(c *C) {
	input := []byte{0x11, 0x12, 0x13, 0x14, 0x15}
	reader := NewTrustworthyReader(bytes.NewBuffer(input))
	first := [2]byte{}

	n, err := reader.Read(first[:])
	c.Assert(first[:], DeepEquals, input[:2])
	c.Assert(n, Equals, len(first))
	c.Assert(err, DeepEquals, nil)

	second := [2]byte{}
	n, err = reader.Read(second[:])
	c.Assert(second[:], DeepEquals, input[2:4])
	c.Assert(n, Equals, len(second))
	c.Assert(err, DeepEquals, nil)

	third := [2]byte{}
	n, err = reader.Read(third[:])
	c.Assert(third[:], DeepEquals, []byte{0x15, 0x00})
	c.Assert(n, Equals, 1)
	c.Assert(err, DeepEquals, io.EOF)
}

func (s *TrustworthyReaderSuite) TestReadExactlyBehavior(c *C) {
	r := newSelfPacedReader([][]byte{
		[]byte{0x01, 0x02},
		[]byte{0x03},
		[]byte{0x04, 0x05},
	})

	reader := NewTrustworthyReader(r)

	into := [4]byte{}
	n, err := reader.Read(into[:])
	c.Assert(into[:], DeepEquals, []byte{0x01, 0x02, 0x03, 0x04})
	c.Assert(n, Equals, len(into))
	c.Assert(err, Equals, nil)
}
