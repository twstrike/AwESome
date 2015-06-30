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

func (s *SBoxTestSuite) TestInvAffineTransformation(c *C) {
	c.Assert(invAffineTrans(byte(0x63)), Equals, byte(0x00))
	c.Assert(invAffineTrans(byte(0xed)), Equals, byte(0xca))
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

func (s *SubBytesTestSuite) TestApplyInvSBox(c *C) {
	c.Assert(applyInvSBox(0x63), Equals, byte(0x00))
	c.Assert(applyInvSBox(0x82), Equals, byte(0x11))
	c.Assert(applyInvSBox(0x93), Equals, byte(0x22))
	c.Assert(applyInvSBox(0xc3), Equals, byte(0x33))
	c.Assert(applyInvSBox(0x1b), Equals, byte(0x44))
	c.Assert(applyInvSBox(0xfc), Equals, byte(0x55))
	c.Assert(applyInvSBox(0x33), Equals, byte(0x66))
	c.Assert(applyInvSBox(0xf5), Equals, byte(0x77))
	c.Assert(applyInvSBox(0xc4), Equals, byte(0x88))
	c.Assert(applyInvSBox(0xee), Equals, byte(0x99))
	c.Assert(applyInvSBox(0xac), Equals, byte(0xaa))
	c.Assert(applyInvSBox(0xea), Equals, byte(0xbb))
	c.Assert(applyInvSBox(0x4b), Equals, byte(0xcc))
	c.Assert(applyInvSBox(0xc1), Equals, byte(0xdd))
	c.Assert(applyInvSBox(0x28), Equals, byte(0xee))
	c.Assert(applyInvSBox(0x16), Equals, byte(0xff))
}
