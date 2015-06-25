package block

import (
	"bytes"
	"encoding/hex"
)

type CBC struct {
	IV []byte
}

func (bm CBC) Encrypt(key, plainText string, blockCipher BlockCipher) string {
	p, _ := hex.DecodeString(plainText)
	k, _ := hex.DecodeString(key)

	ret := bm.encryptBytes(k, p, blockCipher)
	return hex.EncodeToString(ret)
}

func XOR(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("XOR input have different sizes")
	}

	ret := make([]byte, len(a))
	for i := 0; i < len(ret); i++ {
		ret[i] = a[i] ^ b[i]
	}

	return ret
}

func (bm CBC) encryptBytes(key, plain []byte, blockCipher BlockCipher) []byte {
	reader := bytes.NewBuffer(plain)
	cipher := make([]byte, len(plain))
	blockSize := blockCipher.BlockSize() / 8

	firstBlock := XOR(reader.Next(blockSize), bm.IV)
	copy(cipher[:blockSize], blockCipher.Encrypt(key, firstBlock))

	for i := blockSize; reader.Len() > 0; i += blockSize {
		block := XOR(reader.Next(blockSize), cipher[i-blockSize:i])
		copy(cipher[i:], blockCipher.Encrypt(key, block))
	}

	return cipher
}

func (bm CBC) Decrypt(key, cipherText string, blockCipher BlockCipher) string {
	c, _ := hex.DecodeString(cipherText)
	k, _ := hex.DecodeString(key)

	ret := bm.decryptBytes(k, c, blockCipher)
	return hex.EncodeToString(ret)
}

func (bm CBC) decryptBytes(key, cipher []byte, blockCipher BlockCipher) []byte {
	reader := bytes.NewBuffer(cipher)
	plain := make([]byte, len(cipher))
	blockSize := blockCipher.BlockSize() / 8

	previousBlock := reader.Next(blockSize)
	copy(plain[0:blockSize], XOR(blockCipher.Decrypt(key, previousBlock), bm.IV))

	for i := blockSize; reader.Len() > 0; i += blockSize {
		block := reader.Next(blockSize)
		copy(plain[i:], XOR(blockCipher.Decrypt(key, block), previousBlock))
		previousBlock = block
	}

	return plain
}
