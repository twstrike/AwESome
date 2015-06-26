package block

import (
	. "gopkg.in/check.v1"
)

func (s *CBCSuite) Test_xor(c *C) {
	f := func() { xor([]byte{0x01, 0x02}, []byte{0x01, 0x02, 0x03}) }
	c.Assert(f, Panics, "xor: input have different sizes")
}
