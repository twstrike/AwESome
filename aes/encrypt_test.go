package aes

import (
	. "gopkg.in/check.v1"
)

type EncryptSuite struct{}

var _ = Suite(&EncryptSuite{})

func (s *EncryptSuite) Test_stateFromNonZeroes(c *C) {
	// See AES spec, Appendix B – Cipher Example
	res := stateFrom(Block{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
	c.Check(res, DeepEquals, state{
		0x32, 0x88, 0x31, 0xe0,
		0x43, 0x5a, 0x31, 0x37,
		0xf6, 0x30, 0x98, 0x07,
		0xa8, 0x8d, 0xa2, 0x34,
	})
}
