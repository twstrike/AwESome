package aes

import (
	. "gopkg.in/check.v1"
)

type KeysSuite struct{}

var _ = Suite(&KeysSuite{})

func (s *UtilsSuite) Test_aesConfigurationOf128BitKey(c *C) {
	c.Check(Key128{}.aesConfiguration(), Equals, Aes128)
}

func (s *UtilsSuite) Test_aesConfigurationOf192BitKey(c *C) {
	c.Check(Key192{}.aesConfiguration(), Equals, Aes192)
}

func (s *UtilsSuite) Test_aesConfigurationOf256BitKey(c *C) {
	c.Check(Key256{}.aesConfiguration(), Equals, Aes256)
}
