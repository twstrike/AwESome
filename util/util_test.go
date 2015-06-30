package util

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type UtilTestSuite struct{}

var _ = Suite(&UtilTestSuite{})

func (s *UtilTestSuite) Test_rotLeft(c *C) {
	c.Assert(RotLeft(byte(0xaa), 3), Equals, byte(0x55))
}

func (s *UtilTestSuite) Test_rotRight(c *C) {
	c.Assert(RotRight(byte(0x55), 3), Equals, byte(0xaa))
}

func (s *UtilTestSuite) Test_rotLeftUint32(c *C) {
	c.Assert(RotLeftUint32(uint32(0xaabbccdd), 8), Equals, uint32(0xbbccddaa))
}

func (s *UtilTestSuite) Test_rotRightUint32(c *C) {
	c.Assert(RotRightUint32(uint32(0x55443322), 8), Equals, uint32(0x22554433))
}
