package aes

import (
	. "gopkg.in/check.v1"
)

type KeysSuite struct{}

var _ = Suite(&KeysSuite{})

func (s *UtilsSuite) Test_aesConfigurationOf128BitKey(c *C) {
	c.Check(Key128{}.aesConfiguration(), Equals, aes128)
}

func (s *UtilsSuite) Test_aesConfigurationOf192BitKey(c *C) {
	c.Check(Key192{}.aesConfiguration(), Equals, aes192)
}

func (s *UtilsSuite) Test_aesConfigurationOf256BitKey(c *C) {
	c.Check(Key256{}.aesConfiguration(), Equals, aes256)
}

func (s *UtilsSuite) TestParseKeyBytes(c *C) {
	key := []byte{
		0x2b, 0x7e, 0x15, 0x16,
		0x28, 0xae, 0xd2, 0xa6,
		0xab, 0xf7, 0x15, 0x88,
		0x09, 0xcf, 0x4f, 0x3c,
	}
	res := parseKeyFromBytes(key)
	c.Check(res, DeepEquals, Key128{0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x09cf4f3c})
}
