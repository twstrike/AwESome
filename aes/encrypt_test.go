package aes

import (
	. "gopkg.in/check.v1"
	"strings"
)

type EncryptSuite struct{}

var _ = Suite(&EncryptSuite{})

func nospc(s string) HexString {
	return HexString(strings.Replace(s, " ", "", -1))
}

func (s *EncryptSuite) TestParseKeyWithAllZeroes(c *C) {
	res := parseKey("00000000000000000000000000000000")
	c.Check(res, DeepEquals, Key128{0, 0, 0, 0})
}

func (s *EncryptSuite) TestParseKeyWithAOne(c *C) {
	res := parseKey("00000000000000000000000000000001")
	c.Check(res, DeepEquals, Key128{0, 0, 0, 1})
}

func (s *EncryptSuite) TestParseKeyWithSpaces(c *C) {
	res := parseKey(nospc("00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01"))
	c.Check(res, DeepEquals, Key128{0, 0, 0, 1})
}

func (s *EncryptSuite) TestParsePlainTextWithAllZeroes(c *C) {
	res := parsePlainText("00000000000000000000000000000000")
	c.Check(res, DeepEquals, PlainText{0, 0, 0, 0})
}

func (s *EncryptSuite) TestParsePlainTextWithSpaces(c *C) {
	res := parsePlainText("00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00")
	c.Check(res, DeepEquals, PlainText{0, 0, 0, 0})
}

func (s *EncryptSuite) TestParsePlainTextWithAOne(c *C) {
	res := parsePlainText("00000000000000000000000000000001")
	c.Check(res, DeepEquals, PlainText{0, 0, 0, 1})
}

func (s *EncryptSuite) TestToHexStringWithAllZeroes(c *C) {
	res := toHexString(CipherText{0, 0, 0, 0})
	c.Check(res, DeepEquals, HexString("00000000000000000000000000000000"))
}

func (s *EncryptSuite) Test_stateFromAllZeroes(c *C) {
	res := stateFrom(PlainText{0, 0, 0, 0})
	c.Check(res, DeepEquals, state{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func (s *EncryptSuite) Test_stateFromNonZeroes(c *C) {
	// See AES spec, Appendix B – Cipher Example
	res := stateFrom(PlainText{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
	c.Check(res, DeepEquals, state{
		0x32, 0x88, 0x31, 0xe0,
		0x43, 0x5a, 0x31, 0x37,
		0xf6, 0x30, 0x98, 0x07,
		0xa8, 0x8d, 0xa2, 0x34,
	})
}

func (s *EncryptSuite) Test_stateToCipherText(c *C) {
	res := stateToCipherText(state{
		0x32, 0x43, 0xf6, 0xa8,
		0x88, 0x5a, 0x30, 0x8d,
		0x31, 0x31, 0x98, 0xa2,
		0xe0, 0x37, 0x07, 0x34,
	})
	c.Check(res, DeepEquals, CipherText{0x328831e0, 0x435a3137, 0xf6309807, 0xa88da234})
}

func (s *EncryptSuite) TestEncrypt(c *C) {
	plain := PlainText{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734}
	key := Key128{0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x09cf4f3c}
	expected := CipherText{0x3925841d, 0x02dc09fb, 0xdc118597, 0x196a0b32}

	c.Check(expected, DeepEquals, Encrypt(key, plain))
}
