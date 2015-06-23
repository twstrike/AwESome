package aes

import (
	. "gopkg.in/check.v1"
)

type EncryptSuite struct{}

var _ = Suite(&EncryptSuite{})

func (s *EncryptSuite) TestEncrypt128(c *C) {
	//See AES spec, Appendix B – Cipher Example
	plain := Block{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734}
	key := Key128{0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x09cf4f3c}
	expected := Block{0x3925841d, 0x02dc09fb, 0xdc118597, 0x196a0b32}

	c.Check(expected, DeepEquals, Encrypt(key, plain))
}

func (s *EncryptSuite) TestEncrypt192(c *C) {
	plain := Block{0x00112233, 0x44556677, 0x8899aabb, 0xccddeeff}
	key := Key192{0x00010203, 0x04050607, 0x08090a0b, 0x0c0d0e0f, 0x10111213, 0x14151617}
	expected := Block{0xdda97ca4, 0x864cdfe0, 0x6eaf70a0, 0xec0d7191}

	c.Check(expected, DeepEquals, Encrypt(key, plain))
}
