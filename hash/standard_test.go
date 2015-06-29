package hash

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type SHA1Suite struct{}

var _ = Suite(&SHA1Suite{})

var sha1TestCases = TestCase{
	//https://tools.ietf.org/html/rfc3174
	[]string{
		"abc",
		"abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
		"a",
		"0123456701234567012345670123456701234567012345670123456701234567",
	},
	[]string{
		"A9993E364706816ABA3E25717850C26C9CD0D89D",
		"84983E441C3BD26EBAAE4AA1F95129E5E54670F1",
		"34AA973CD4C4DAA4F61EEB2BDBAD27316534016F",
		"DEA356A2CDDD90C7A7ECEDC5EBB563934F460452",
	},
}

func (s *SHA1Suite) TestSHA1(c *C) {
	c.Skip("not implemented")
	sha1 := SHA1{}
	testOnSum(sha1, sha1TestCases, c)
}
