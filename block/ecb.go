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
	cipher := make([]byte, len(plain))
	blockSize := blockCipher.BlockSize() / 8

	for i := 0; reader.Len() > 0; i += blockSize {
		block := reader.Next(blockSize)
		copy(cipher[i:], blockCipher.Encrypt(key, block))
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
	plain := make([]byte, len(cipher))
	blockSize := blockCipher.BlockSize() / 8

	for i := 0; reader.Len() > 0; i += blockSize {
		block := reader.Next(blockCipher.BlockSize() / 8)
		copy(plain[i:], blockCipher.Decrypt(key, block))
	}

	return plain
}
