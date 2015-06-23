package aes

import (
	. "gopkg.in/check.v1"
)

type ShiftRowsSuite struct{}

var _ = Suite(&ShiftRowsSuite{})

func (s *ShiftRowsSuite) TestShiftRowsWithSimpleData(c *C) {
	input := state{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
	}

	expected := state{
		0, 1, 2, 3,
		5, 6, 7, 4,
		10, 11, 8, 9,
		15, 12, 13, 14,
	}

	c.Check(shiftRows(input), DeepEquals, expected)
}

func (s *ShiftRowsSuite) TestInvShiftRowsWithSimpleData(c *C) {
	input := state{
		0, 1, 2, 3,
		5, 6, 7, 4,
		10, 11, 8, 9,
		15, 12, 13, 14,
	}

	expected := state{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
	}

	c.Check(invShiftRows(input), DeepEquals, expected)
}
