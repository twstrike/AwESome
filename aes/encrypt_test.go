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

func (s *EncryptSuite) TestEncryptBytes128(c *C) {
	//See AES spec, Appendix B – Cipher Example
	plain := []byte{
		0x32, 0x43, 0xf6, 0xa8,
		0x88, 0x5a, 0x30, 0x8d,
		0x31, 0x31, 0x98, 0xa2,
		0xe0, 0x37, 0x07, 0x34,
	}
	key := []byte{
		0x2b, 0x7e, 0x15, 0x16,
		0x28, 0xae, 0xd2, 0xa6,
		0xab, 0xf7, 0x15, 0x88,
		0x09, 0xcf, 0x4f, 0x3c,
	}
	expected := []byte{
		0x39, 0x25, 0x84, 0x1d,
		0x02, 0xdc, 0x09, 0xfb,
		0xdc, 0x11, 0x85, 0x97,
		0x19, 0x6a, 0x0b, 0x32,
	}

	c.Check(expected, DeepEquals, EncryptBytes(key, plain))
}
