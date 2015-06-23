package aes

import (
	. "gopkg.in/check.v1"
)

type BlockSuite struct{}

var _ = Suite(&BlockSuite{})

func (s *BlockSuite) Test_stateFromAllZeroes(c *C) {
	res := Block{0, 0, 0, 0}.toState()
	c.Check(res, DeepEquals, state{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func (s *BlockSuite) Test_stateFromNonZeroes(c *C) {
	// See AES spec, Appendix B – Cipher Example
  res := Block{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734}.toState()
	c.Check(res, DeepEquals, state{
		0x32, 0x88, 0x31, 0xe0,
		0x43, 0x5a, 0x31, 0x37,
		0xf6, 0x30, 0x98, 0x07,
		0xa8, 0x8d, 0xa2, 0x34,
	})
}

func (s *BlockSuite) TestToHexStringWithAllZeroes(c *C) {
	res := Block{0, 0, 0, 0}.toHexString()
	c.Check(res, DeepEquals, HexString("00000000000000000000000000000000"))
}

func (s *BlockSuite) TestWordToHexStringWith4Words(c *C) {
	res := Block{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734}.toHexString()
	c.Check(res, DeepEquals, HexString("3243f6a8885a308d313198a2e0370734"))
}
