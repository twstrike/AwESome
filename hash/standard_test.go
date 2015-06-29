package hash

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type SHA1Suite struct{}

var _ = Suite(&SHA1Suite{})

//https://tools.ietf.org/html/rfc3174
var sha1TestCases = []TestCase{
	TestCase{
		"abc",
		"A9993E364706816ABA3E25717850C26C9CD0D89D",
		1,
	},
	TestCase{
		"abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
		"84983E441C3BD26EBAAE4AA1F95129E5E54670F1",
		1,
	},
	TestCase{
		"a",
		"34AA973CD4C4DAA4F61EEB2BDBAD27316534016F",
		1000000,
	},
	TestCase{
		"0123456701234567012345670123456701234567012345670123456701234567",
		"DEA356A2CDDD90C7A7ECEDC5EBB563934F460452",
		10,
	},
}

func (s *SHA1Suite) TestSHA1(c *C) {
	c.Skip("not implemented")
	sha1 := SHA1{}
	for _, tc := range sha1TestCases {
		testOnSum(sha1, tc, c)
	}
}
