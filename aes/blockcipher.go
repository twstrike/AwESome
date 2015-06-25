package aes

var BlockCipher = blockCipher{}

type blockCipher struct{}

func (b blockCipher) BlockSize() int {
	return 128
}

func (b blockCipher) Encrypt(plainText, key []byte) []byte {
	return EncryptBytes(key, plainText)
}

func (b blockCipher) Decrypt(cipherText, key []byte) []byte {
	return DecryptBytes(key, cipherText)
}
