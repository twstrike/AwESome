package rijndael

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type RijndaelTestSuite struct{}

var _ = Suite(&RijndaelTestSuite{})

func (s *RijndaelTestSuite) TestMultiplication(c *C) {
	var a, b byte
	a = 0x57
	b = 0x83

	c.Assert(multiplication(a, b), Equals, uint16(0x2b79))
}

func (s *RijndaelTestSuite) TestModulo(c *C) {
	var a, b uint16
	a = 0x2b79
	b = 0x011b

	c.Assert(modulo(a, b), Equals, byte(0xc1))
}

func (s *RijndaelTestSuite) TestNbits(c *C) {
	var a, b uint16
	a = 0x2b79
	b = 0x011b

	c.Assert(nbits(a), Equals, uint(14))
	c.Assert(nbits(b), Equals, uint(9))
}

func (s *RijndaelTestSuite) TestMultiplicativeInverse(c *C) {
	subject := byte(0x53)
	expected := byte(0xCA)

	c.Assert(Inv(subject), Equals, expected)
	c.Assert(Inv(0), Equals, byte(0))
}
