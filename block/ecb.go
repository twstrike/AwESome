package block

import (
	"bytes"
	"encoding/hex"
)

type BlockCipher interface {
	BlockSize() int
	Encrypt(block, key string) string
	Decrypt(block, key string) string
}

func (bm ECB) Encrypt(plainText, key string, cipher BlockCipher) string {
	data, _ := hex.DecodeString(plainText)
	reader := bytes.NewBuffer(data)
	cipherText := ""

	for reader.Len() > 0 {
		block := reader.Next(cipher.BlockSize() / 8)
		cipherText += cipher.Encrypt(hex.EncodeToString(block), key)
	}

	return cipherText
}

func (bm ECB) Decrypt(cipherText, key string, cipher BlockCipher) string {
	data, _ := hex.DecodeString(cipherText)
	reader := bytes.NewBuffer(data)
	plainText := ""

	for reader.Len() > 0 {
		block := reader.Next(cipher.BlockSize() / 8)
		plainText += cipher.Decrypt(hex.EncodeToString(block), key)
	}

	return plainText
}
