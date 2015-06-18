package block_test

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type BlockSuite struct{}

var _ = Suite(&BlockSuite{})

func (s *BlockSuite) TestHelloWorld(c *C) {
	c.Assert(42, Equals, 42)
	c.Check("42", Equals, "42")
	c.Check(42, Equals, 42)
}
