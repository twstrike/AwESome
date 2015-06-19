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
	res := stateFrom(PlainText{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
	c.Check(res, DeepEquals, state{
		0x32, 0x43, 0xf6, 0xa8,
		0x88, 0x5a, 0x30, 0x8d,
		0x31, 0x31, 0x98, 0xa2,
		0xe0, 0x37, 0x07, 0x34,
	})
}
