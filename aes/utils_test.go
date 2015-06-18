package aes

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TesthexStringToWordEmptyHexString(c *C) {
	var result [0]word
	hexStringToWord("", &result)
	c.Check(result, DeepEquals, [0]word{})
}

func (s *UtilsSuite) TesthexStringToWordWith4Words(c *C) {
	var result [4]word
	hexStringToWord("3243f6a8885a308d313198a2e0370734", &result)
	c.Check(result, DeepEquals, [4]word{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
}

func (s *UtilsSuite) TesthexStringToWordWith2Words(c *C) {
	var result [2]word
	hexStringToWord("313198a2e0370734", &result)
	c.Check(result, DeepEquals, [2]word{0x313198a2, 0xe0370734})
}
