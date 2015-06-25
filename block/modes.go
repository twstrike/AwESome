package block

type BlockMode interface {
	Encrypt(plain, key []byte, bc BlockCipher) []byte
	Decrypt(plain, key []byte, bc BlockCipher) []byte
}
