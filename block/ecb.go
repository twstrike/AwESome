package block

import (
	"bytes"
	"encoding/hex"
)

func (bm ECB) Encrypt(key, plainText string, blockCipher BlockCipher) string {
	p, _ := hex.DecodeString(plainText)
	k, _ := hex.DecodeString(key)

	ret := bm.encryptBytes(k, p, blockCipher)
	return hex.EncodeToString(ret)
}

func (bm ECB) encryptBytes(key, plain []byte, blockCipher BlockCipher) []byte {
	reader := bytes.NewBuffer(plain)
	cipher := make([]byte, 0, len(plain))

	for reader.Len() > 0 {
		block := reader.Next(blockCipher.BlockSize() / 8)
		cipher = append(cipher, blockCipher.Encrypt(key, block)...)
	}

	return cipher
}

func (bm ECB) Decrypt(key, cipherText string, blockCipher BlockCipher) string {
	c, _ := hex.DecodeString(cipherText)
	k, _ := hex.DecodeString(key)

	ret := bm.decryptBytes(k, c, blockCipher)
	return hex.EncodeToString(ret)
}

func (bm ECB) decryptBytes(key, cipher []byte, blockCipher BlockCipher) []byte {
	reader := bytes.NewBuffer(cipher)
	plain := make([]byte, 0, len(cipher))

	for reader.Len() > 0 {
		block := reader.Next(blockCipher.BlockSize() / 8)
		plain = append(plain, blockCipher.Decrypt(key, block)...)
	}

	return plain
}
