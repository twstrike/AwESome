package hash

import (
	"encoding/hex"
	. "gopkg.in/check.v1"
)

type TestCase struct {
	inputs  []string
	results []string
}

func testOnSum(hash Hash, tc TestCase, c *C) {
	hash.Init()
	inputs := tc.inputs
	results := tc.results
	for i := range inputs {
		b := []byte(inputs[i])
		r, _ := hex.DecodeString(results[i])
		c.Assert(hash.Sum(b), DeepEquals, r)
	}
}
