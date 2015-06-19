package aes

import (
	. "gopkg.in/check.v1"
)

type GaloisTestSuite struct{}

var _ = Suite(&GaloisTestSuite{})

func (s *GaloisTestSuite) TestMultiplication(c *C) {
	var a, b byte
	a = 0x57
	b = 0x83

	c.Assert(multiplication(a, b), Equals, uint16(0x2b79))
}

func (s *GaloisTestSuite) TestModulo(c *C) {
	var a, b uint16
	a = 0x2b79
	b = 0x011b

	c.Assert(modulo(a, b), Equals, byte(0xc1))
}

func (s *GaloisTestSuite) TestNbits(c *C) {
	var a, b uint16
	a = 0x2b79
	b = 0x011b

	c.Assert(nbits(a), Equals, uint(14))
	c.Assert(nbits(b), Equals, uint(9))
}
