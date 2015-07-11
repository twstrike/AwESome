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
		c.Assert(invSboxTable[b], Equals, calculateInvSBox(b))

		c.Assert(invSboxTable[sboxTable[b]], Equals, b)
		c.Assert(sboxTable[invSboxTable[b]], Equals, b)
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
