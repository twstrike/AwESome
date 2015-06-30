package block

import (
	"bytes"
)

type ECB struct{}

func (bm ECB) Encrypt(key, plain []byte, blockCipher Cipher) []byte {
	reader := bytes.NewBuffer(plain)
	cipher := make([]byte, len(plain))
	blockSize := blockCipher.BlockSize() / 8

	for i := 0; reader.Len() > 0; i += blockSize {
		block := reader.Next(blockSize)
		copy(cipher[i:], blockCipher.Encrypt(key, block))
	}

	return cipher
}

func (bm ECB) Decrypt(key, cipher []byte, blockCipher Cipher) []byte {
	reader := bytes.NewBuffer(cipher)
	plain := make([]byte, len(cipher))
	blockSize := blockCipher.BlockSize() / 8

	for i := 0; reader.Len() > 0; i += blockSize {
		block := reader.Next(blockCipher.BlockSize() / 8)
		copy(plain[i:], blockCipher.Decrypt(key, block))
	}

	return plain
}
