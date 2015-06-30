package block

import (
	"bytes"
)

type CBC struct {
	IV []byte
}

func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: input have different sizes")
	}

	ret := make([]byte, len(a))
	for i := 0; i < len(ret); i++ {
		ret[i] = a[i] ^ b[i]
	}

	return ret
}

func (bm CBC) Encrypt(key, plain []byte, blockCipher Cipher) []byte {
	reader := bytes.NewBuffer(plain)
	cipher := make([]byte, len(plain))
	blockSize := blockCipher.BlockSize() / 8

	firstBlock := xor(reader.Next(blockSize), bm.IV)
	copy(cipher[:blockSize], blockCipher.Encrypt(key, firstBlock))

	for i := blockSize; reader.Len() > 0; i += blockSize {
		block := xor(reader.Next(blockSize), cipher[i-blockSize:i])
		copy(cipher[i:], blockCipher.Encrypt(key, block))
	}

	return cipher
}

func (bm CBC) Decrypt(key, cipher []byte, blockCipher Cipher) []byte {
	reader := bytes.NewBuffer(cipher)
	plain := make([]byte, len(cipher))
	blockSize := blockCipher.BlockSize() / 8

	previousBlock := reader.Next(blockSize)
	copy(plain[0:blockSize], xor(blockCipher.Decrypt(key, previousBlock), bm.IV))

	for i := blockSize; reader.Len() > 0; i += blockSize {
		block := reader.Next(blockSize)
		copy(plain[i:], xor(blockCipher.Decrypt(key, block), previousBlock))
		previousBlock = block
	}

	return plain
}
