package aes

import (
	. "gopkg.in/check.v1"
)

type SBoxTestSuite struct{}

var _ = Suite(&SBoxTestSuite{})

func (s *SBoxTestSuite) TestAffineTransformation(c *C) {
	c.Assert(affineTrans(byte(0x00)), Equals, byte(0x63))
	c.Assert(affineTrans(byte(0xca)), Equals, byte(0xed))
}

func (s *SubBytesTestSuite) TestApplySBox(c *C) {
	c.Assert(applySBox(0x00), Equals, byte(0x63))
	c.Assert(applySBox(0x11), Equals, byte(0x82))
	c.Assert(applySBox(0x22), Equals, byte(0x93))
	c.Assert(applySBox(0x33), Equals, byte(0xc3))
	c.Assert(applySBox(0x44), Equals, byte(0x1b))
	c.Assert(applySBox(0x55), Equals, byte(0xfc))
	c.Assert(applySBox(0x66), Equals, byte(0x33))
	c.Assert(applySBox(0x77), Equals, byte(0xf5))
	c.Assert(applySBox(0x88), Equals, byte(0xc4))
	c.Assert(applySBox(0x99), Equals, byte(0xee))
	c.Assert(applySBox(0xaa), Equals, byte(0xac))
	c.Assert(applySBox(0xbb), Equals, byte(0xea))
	c.Assert(applySBox(0xcc), Equals, byte(0x4b))
	c.Assert(applySBox(0xdd), Equals, byte(0xc1))
	c.Assert(applySBox(0xee), Equals, byte(0x28))
	c.Assert(applySBox(0xff), Equals, byte(0x16))
}
