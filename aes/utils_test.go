package aes

import (
	. "gopkg.in/check.v1"
	"strings"
)

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TesthexStringToWordEmptyHexString(c *C) {
	var result [0]word
	hexStringToWord("", &result)
	c.Check(result, DeepEquals, [0]word{})
}

func (s *UtilsSuite) TesthexStringToWordWith4Words(c *C) {
	var result [4]word
	hexStringToWord("3243f6a8885a308d313198a2e0370734", &result)
	c.Check(result, DeepEquals, [4]word{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
}

func (s *UtilsSuite) TesthexStringToWordWith2Words(c *C) {
	var result [2]word
	hexStringToWord("313198a2e0370734", &result)
	c.Check(result, DeepEquals, [2]word{0x313198a2, 0xe0370734})
}

func nospc(s string) HexString {
	return HexString(strings.Replace(s, " ", "", -1))
}

func (s *UtilsSuite) TestParseKeyWithAllZeroes(c *C) {
	res := parseKey("00000000000000000000000000000000")
	c.Check(res, DeepEquals, Key128{0, 0, 0, 0})
}

func (s *UtilsSuite) TestParseKeyWithAOne(c *C) {
	res := parseKey("00000000000000000000000000000001")
	c.Check(res, DeepEquals, Key128{0, 0, 0, 1})
}

func (s *UtilsSuite) TestParseKeyWithSpaces(c *C) {
	res := parseKey(nospc("00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01"))
	c.Check(res, DeepEquals, Key128{0, 0, 0, 1})
}

func (s *UtilsSuite) TestParsePlainTextWithAllZeroes(c *C) {
	res := HexString("00000000000000000000000000000000").toBlock()
	c.Check(res, DeepEquals, Block{0, 0, 0, 0})
}

func (s *UtilsSuite) TestParsePlainTextWithAOne(c *C) {
	res := HexString("00000000000000000000000000000001").toBlock()
	c.Check(res, DeepEquals, Block{0, 0, 0, 1})
}

func (s *UtilsSuite) TestParsePlainTextWithValues(c *C) {
	res := HexString("3243f6a8885a308d313198a2e0370734").toBlock()
	c.Check(res, DeepEquals, Block{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
}
