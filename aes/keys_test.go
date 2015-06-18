package aes

import (
	. "gopkg.in/check.v1"
)

type KeysSuite struct{}

var _ = Suite(&KeysSuite{})

func (s *UtilsSuite) Test_keyLengthOf128BitKey(c *C) {
	c.Check(Key128{}.keyLength(), Equals, 128)
}

func (s *UtilsSuite) Test_keyLengthOf192BitKey(c *C) {
	c.Check(Key192{}.keyLength(), Equals, 192)
}

func (s *UtilsSuite) Test_keyLengthOf256BitKey(c *C) {
	c.Check(Key256{}.keyLength(), Equals, 256)
}
