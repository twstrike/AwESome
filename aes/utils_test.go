package aes

import (
	"encoding/hex"

	. "gopkg.in/check.v1"
)

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestZeroBytesToWord(c *C) {
	var result [0]word
	bytesToWord([]byte{}, &result)
	c.Check(result, DeepEquals, [0]word{})
}

func (s *UtilsSuite) TestBytesToFourWords(c *C) {
	var result [4]word
	bytes, _ := hex.DecodeString("3243f6a8885a308d313198a2e0370734")

	bytesToWord(bytes, &result)
	c.Check(result, DeepEquals, [4]word{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
}

func (s *UtilsSuite) TestWordsToBytes(c *C) {
	var result = wordsToBytes([4]word{0x3243f6a8, 0x885a308d, 0x313198a2, 0xe0370734})
	c.Check(result, DeepEquals, []byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x7, 0x34})
}
