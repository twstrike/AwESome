package block

import (
	"bytes"
	"encoding/hex"
)

func (bm ECB) Encrypt(plainText, key string, blockCipher BlockCipher) string {
	p, _ := hex.DecodeString(plainText)
	k, _ := hex.DecodeString(key)

	ret := bm.encryptBytes(p, k, blockCipher)
	return hex.EncodeToString(ret)
}

func (bm ECB) encryptBytes(plain, key []byte, blockCipher BlockCipher) []byte {
	reader := bytes.NewBuffer(plain)
	cipher := make([]byte, 0, len(plain))

	for reader.Len() > 0 {
		block := reader.Next(blockCipher.BlockSize() / 8)
		cipher = append(cipher, blockCipher.Encrypt(block, key)...)
	}

	return cipher
}

func (bm ECB) Decrypt(cipherText, key string, blockCipher BlockCipher) string {
	c, _ := hex.DecodeString(cipherText)
	k, _ := hex.DecodeString(key)

	ret := bm.decryptBytes(c, k, blockCipher)
	return hex.EncodeToString(ret)
}

func (bm ECB) decryptBytes(cipher, key []byte, blockCipher BlockCipher) []byte {
	reader := bytes.NewBuffer(cipher)
	plain := make([]byte, 0, len(cipher))

	for reader.Len() > 0 {
		block := reader.Next(blockCipher.BlockSize() / 8)
		plain = append(plain, blockCipher.Decrypt(block, key)...)
	}

	return plain
}
