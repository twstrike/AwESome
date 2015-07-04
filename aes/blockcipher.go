package aes

// Represent a Block Cipher that can encrypt and decrypt a fixed size byte block
var BlockCipher = blockCipher{}

type blockCipher struct{}

func (b blockCipher) BlockSize() int {
	return 128
}

func (b blockCipher) Encrypt(key, plainText []byte) []byte {
	return EncryptBytes(key, plainText)
}

func (b blockCipher) Decrypt(key, cipherText []byte) []byte {
	return DecryptBytes(key, cipherText)
}
