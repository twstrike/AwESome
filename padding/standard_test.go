package padding

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type PaddingSuite struct{}

var _ = Suite(&PaddingSuite{})

func (s *PaddingSuite) TestPKCS7PaddingSingleBlockAndLesserThanBlockSize(c *C) {
	data := []byte{0xaa, 0xbb, 0xcc}
	result := PKCS7(data, 8)
	c.Assert([]byte{0xaa, 0xbb, 0xcc, 0x5, 0x5, 0x5, 0x5, 0x5}, DeepEquals, result)
}

func (s *PaddingSuite) TestPKCS7PaddingMultipleBlocksAndLesserThanBlockSize(c *C) {
	data := []byte{0xaa, 0xbb, 0xcc}
	result := PKCS7(data, 2)
	c.Assert([]byte{0xaa, 0xbb, 0xcc, 0x01}, DeepEquals, result)
}

func (s *PaddingSuite) TestPKCS7PaddingCreatesAdditionalBlockWhenInputIsMultipleOfBlockSize(c *C) {
	data := []byte{0xaa, 0xbb, 0xcc, 0xdd}
	result := PKCS7(data, 2)
	c.Assert([]byte{0xaa, 0xbb, 0xcc, 0xdd, 0x02, 0x02}, DeepEquals, result)
}

func (s *PaddingSuite) TestUnpaddingPKCS7(c *C) {
	data := []byte{0xaa, 0xbb, 0xcc, 0x5, 0x5, 0x5, 0x5, 0x5}
	result := UndoPKCS7(data)
	c.Assert([]byte{0xaa, 0xbb, 0xcc}, DeepEquals, result)
}

func (s *PaddingSuite) TestPKCS7UnpaddingAdditionalBlock(c *C) {
	data := []byte{0xaa, 0xbb, 0xcc, 0xdd, 0x02, 0x02}
	result := UndoPKCS7(data)
	c.Assert([]byte{0xaa, 0xbb, 0xcc, 0xdd}, DeepEquals, result)
}
