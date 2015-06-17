package block

import (
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ECBSuite struct{}
var _ = Suite(&ECBSuite{})

type CBCSuite struct{}
var _ = Suite(&CBCSuite{})

type CFB1Suite struct{}
var _ = Suite(&CFB1Suite{})

type CFB8Suite struct{}
var _ = Suite(&CFB8Suite{})

type CFB128Suite struct{}
var _ = Suite(&CFB128Suite{})

type OFBSuite struct{}
var _ = Suite(&OFBSuite{})

type CTRSuite struct{}
var _ = Suite(&CTRSuite{})

// Data from: http://csrc.nist.gov/publications/nistpubs/800-38a/sp800-38a.pdf

type Block struct {
	plain, cipher string
}

type TestVector struct {
	key string
	iv  string
	blocks []Block
}

var ecbAES128Vector = TestVector{
	"2b7e151628aed2a6abf7158809cf4f3c",
	"",
	[]Block{
		Block{
			"6bc1bee22e409f96e93d7e117393172a",
			"3ad77bb40d7a3660a89ecaf32466ef97",
		},
		Block{
			"ae2d8a571e03ac9c9eb76fac45af8e51",
			"f5d3d58503b9699de785895a96fdbaaf",
		},
		Block{
			"30c81c46a35ce411e5fbc1191a0a52ef",
			"43b1cd7f598ece23881b00e3ed030688",
		},
		Block{
			"f69f2445df4f9b17ad2b417be66c3710",
			"7b0c785e27e8ad3f8223207104725dd4",
		},
	},
}

var ecbAES192Vector = TestVector{
	"8e73b0f7da0e6452c810f32b809079e562f8ead2522c6b7b",
	"",
	[]Block{
		Block{
			"6bc1bee22e409f96e93d7e117393172a",
			"bd334f1d6e45f25ff712a214571fa5cc",
		},
		Block{
			"ae2d8a571e03ac9c9eb76fac45af8e51",
			"974104846d0ad3ad7734ecb3ecee4eef",
		},
		Block{
			"30c81c46a35ce411e5fbc1191a0a52ef",
			"ef7afd2270e2e60adce0ba2face6444e",
		},
		Block{
			"f69f2445df4f9b17ad2b417be66c3710",
			"9a4b41ba738d6c72fb16691603c18e0e",
		},
	},
}

var cbcAES128Vector = TestVector{
	"2b7e151628aed2a6abf7158809cf4f3c",
	"000102030405060708090a0b0c0d0e0f",
	[]Block{
		Block{
			"6bc1bee22e409f96e93d7e117393172a",
			"7649abac8119b246cee98e9b12e9197d",
		},
		Block{
			"ae2d8a571e03ac9c9eb76fac45af8e51",
			"5086cb9b507219ee95db113a917678b2",
		},
		Block{
			"30c81c46a35ce411e5fbc1191a0a52ef",
			"73bed6b8e3c1743b7116e69e22229516",
		},
		Block{
			"f69f2445df4f9b17ad2b417be66c3710",
			"3ff1caa1681fac09120eca307586e1a7",
		},
	},
}
