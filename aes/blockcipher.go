package aes

var BlockCipher = blockCipher{}

type blockCipher struct{}

func (b blockCipher) BlockSize() int {
	return 128
}

func (b blockCipher) Encrypt(plainText, key string) string {
	return string(EncryptHex(HexString(key), HexString(plainText)))
}

func (b blockCipher) Decrypt(cipherText, key string) string {
	return string(DecryptHex(HexString(key), HexString(cipherText)))
}
