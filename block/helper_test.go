package block

import (
	. "gopkg.in/check.v1"
)

type Block struct {
	plain, cipher string
}

type TestVector struct {
	key    string
	iv     string
	blocks []Block
}

func testOnEncrypt(bm BlockMode, bc BlockCipher, tv TestVector, c *C) {
	plain := ""
	cipher := ""
	for b := range tv.blocks {
		plain += tv.blocks[b].plain
		cipher += tv.blocks[b].cipher
	}
	result := bm.Encrypt(tv.key, plain, bc)
	c.Assert(result, Equals, cipher)
}

func testOnDecrypt(bm BlockMode, bc BlockCipher, tv TestVector, c *C) {
	plain := ""
	cipher := ""
	for b := range tv.blocks {
		plain += tv.blocks[b].plain
		cipher += tv.blocks[b].cipher
	}
	result := bm.Decrypt(tv.key, cipher, bc)
	c.Assert(result, Equals, plain)
}
