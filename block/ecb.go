package block

import (
	"bytes"
)

// ECB contains the information necessary to run the ECB block mode
type ECB struct{}

// Encrypt encrypts the plain text with the key using the provided block cipher and ECB block mode
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

// Decrypt decrypts the cipher text with the key using the provided block cipher and ECB block mode
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
