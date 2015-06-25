package block

import (
	"encoding/hex"
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

	key, _ := hex.DecodeString(tv.key)
	p, _ := hex.DecodeString(plain)
	result := hex.EncodeToString(bm.Encrypt(key, p, bc))
	c.Assert(result, Equals, cipher)
}

func testOnDecrypt(bm BlockMode, bc BlockCipher, tv TestVector, c *C) {
	plain := ""
	cipher := ""
	for b := range tv.blocks {
		plain += tv.blocks[b].plain
		cipher += tv.blocks[b].cipher
	}
	key, _ := hex.DecodeString(tv.key)
	ct, _ := hex.DecodeString(cipher)
	result := hex.EncodeToString(bm.Decrypt(key, ct, bc))
	c.Assert(result, Equals, plain)
}
