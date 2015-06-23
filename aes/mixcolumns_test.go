package aes

import (
	. "gopkg.in/check.v1"
)

type MixColumnsSuite struct{}

var _ = Suite(&MixColumnsSuite{})

func (s *MixColumnsSuite) Test_mixOneColumnTestVector1(c *C) {
	before := state{
		0xDB, 0, 0, 0,
		0x13, 0, 0, 0,
		0x53, 0, 0, 0,
		0x45, 0, 0, 0,
	}
	expected := state{
		0x8E, 0, 0, 0,
		0x4D, 0, 0, 0,
		0xA1, 0, 0, 0,
		0xBC, 0, 0, 0,
	}

	c.Check(mixOneColumn(before, 0), DeepEquals, expected)
}

func (s *MixColumnsSuite) Test_mixOneColumnTestVector2(c *C) {
	before := state{
		0xF2, 0, 0, 0,
		0x0A, 0, 0, 0,
		0x22, 0, 0, 0,
		0x5C, 0, 0, 0,
	}
	expected := state{
		0x9F, 0, 0, 0,
		0xDC, 0, 0, 0,
		0x58, 0, 0, 0,
		0x9D, 0, 0, 0,
	}

	c.Check(mixOneColumn(before, 0), DeepEquals, expected)
}

func (s *MixColumnsSuite) Test_mixOneColumnTestVector3(c *C) {
	before := state{
		0x01, 0, 0, 0,
		0x01, 0, 0, 0,
		0x01, 0, 0, 0,
		0x01, 0, 0, 0,
	}
	expected := state{
		0x01, 0, 0, 0,
		0x01, 0, 0, 0,
		0x01, 0, 0, 0,
		0x01, 0, 0, 0,
	}

	c.Check(mixOneColumn(before, 0), DeepEquals, expected)
}

func (s *MixColumnsSuite) Test_mixOneColumnTestVector4(c *C) {
	before := state{
		0, 0xC6, 0, 0,
		0, 0xC6, 0, 0,
		0, 0xC6, 0, 0,
		0, 0xC6, 0, 0,
	}
	expected := state{
		0, 0xC6, 0, 0,
		0, 0xC6, 0, 0,
		0, 0xC6, 0, 0,
		0, 0xC6, 0, 0,
	}

	c.Check(mixOneColumn(before, 1), DeepEquals, expected)
}

func (s *MixColumnsSuite) Test_mixOneColumnTestVector5(c *C) {
	before := state{
		0, 0, 0xD4, 0,
		0, 0, 0xD4, 0,
		0, 0, 0xD4, 0,
		0, 0, 0xD5, 0,
	}
	expected := state{
		0, 0, 0xD5, 0,
		0, 0, 0xD5, 0,
		0, 0, 0xD7, 0,
		0, 0, 0xD6, 0,
	}

	c.Check(mixOneColumn(before, 2), DeepEquals, expected)
}

func (s *MixColumnsSuite) Test_mixOneColumnTestVector6(c *C) {
	before := state{
		0, 0, 0x2D, 0,
		0, 0, 0x26, 0,
		0, 0, 0x31, 0,
		0, 0, 0x4C, 0,
	}
	expected := state{
		0, 0, 0x4D, 0,
		0, 0, 0x7E, 0,
		0, 0, 0xBD, 0,
		0, 0, 0xF8, 0,
	}

	c.Check(mixOneColumn(before, 2), DeepEquals, expected)
}

func (s *MixColumnsSuite) Test_mixColumnTestVector(c *C) {
	before := state{
		0xDB, 0xC6, 0x2D, 0x01,
		0x13, 0xC6, 0x26, 0x01,
		0x53, 0xC6, 0x31, 0x01,
		0x45, 0xC6, 0x4C, 0x01,
	}
	expected := state{
		0x8E, 0xC6, 0x4D, 0x01,
		0x4D, 0xC6, 0x7E, 0x01,
		0xA1, 0xC6, 0xBD, 0x01,
		0xBC, 0xC6, 0xF8, 0x01,
	}

	c.Check(mixColumns(before), DeepEquals, expected)
}

func (s *MixColumnsSuite) Test_invMixColumnTestVector(c *C) {
	before := state{
		0x8E, 0xC6, 0x4D, 0x01,
		0x4D, 0xC6, 0x7E, 0x01,
		0xA1, 0xC6, 0xBD, 0x01,
		0xBC, 0xC6, 0xF8, 0x01,
	}
	expected := state{
		0xDB, 0xC6, 0x2D, 0x01,
		0x13, 0xC6, 0x26, 0x01,
		0x53, 0xC6, 0x31, 0x01,
		0x45, 0xC6, 0x4C, 0x01,
	}

	c.Check(invMixColumns(before), DeepEquals, expected)
}
