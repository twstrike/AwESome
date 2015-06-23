package aes

import (
	. "gopkg.in/check.v1"
)

type DecryptSuite struct{}

var _ = Suite(&DecryptSuite{})

func (s *DecryptSuite) TestDecrypt(c *C) {
	//See AES spec, Appendix B – Cipher Example
	cipher := Block{0x3925841d, 0x02dc09fb, 0xdc118597, 0x196a0b32}
	key := Key128{0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x09cf4f3c}
	expected := Block{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734}

	c.Check(expected, DeepEquals, Decrypt(key, cipher))
}
