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
