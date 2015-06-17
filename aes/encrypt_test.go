package aes

import (
	. "gopkg.in/check.v1"
)

type EncryptSuite struct{}
var _ = Suite(&EncryptSuite{})

func (s *EncryptSuite) TestParseKeyWithAllZeroes(c *C) {
	res := parseKey("00000000000000000000000000000000")
	c.Check(res, DeepEquals, Key128{0,0,0,0})
}

func (s *EncryptSuite) TestParseKeyWithAOne(c *C) {
	res := parseKey("00000000000000000000000000000001")
	c.Check(res, DeepEquals, Key128{0,0,0,1})
}

func (s *EncryptSuite) TestParseKeyWithSpaces(c *C) {
	res := parseKey("00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01")
	c.Check(res, DeepEquals, Key128{0,0,0,1})
}
