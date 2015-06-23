package aes

import (
	. "gopkg.in/check.v1"
)

type StateSuite struct{}

var _ = Suite(&StateSuite{})

func (s *StateSuite) Test_toBlock(c *C) {
	st := state{
		0x32, 0x43, 0xf6, 0xa8,
		0x88, 0x5a, 0x30, 0x8d,
		0x31, 0x31, 0x98, 0xa2,
		0xe0, 0x37, 0x07, 0x34,
	}
	res := st.toBlock()
	c.Check(res, DeepEquals, Block{0x328831e0, 0x435a3137, 0xf6309807, 0xa88da234})
}
