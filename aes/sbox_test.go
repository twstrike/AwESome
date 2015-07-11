package aes

import (
	. "gopkg.in/check.v1"
)

type SBoxTestSuite struct{}

var _ = Suite(&SBoxTestSuite{})

func (*SBoxTestSuite) TestSboxTable(c *C) {
	for i := 0; i <= 256; i++ {
		b := byte(i)
		c.Assert(sboxTable[b], Equals, calculateSBox(b))
	}
}

func (s *SBoxTestSuite) TestAffineTransformation(c *C) {
	c.Assert(affineTrans(byte(0x00)), Equals, byte(0x63))
	c.Assert(affineTrans(byte(0xca)), Equals, byte(0xed))
}

func (s *SBoxTestSuite) TestInvAffineTransformation(c *C) {
	c.Assert(invAffineTrans(byte(0x63)), Equals, byte(0x00))
	c.Assert(invAffineTrans(byte(0xed)), Equals, byte(0xca))
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
