package hash

import (
	"bytes"
	"encoding/hex"

	. "gopkg.in/check.v1"
)

type TestCase struct {
	input  string
	result string
	repeat int
}

func testOnSum(hash Hash, tc TestCase, c *C) {
	current := []byte(tc.input)
	for i := 0; i < tc.repeat; i++ {
		current = hash.Sum(bytes.NewBuffer(current))
	}

	r, _ := hex.DecodeString(tc.result)
	c.Assert(current, DeepEquals, r)
}
