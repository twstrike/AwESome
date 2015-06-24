package block

type ECB struct{}

type BlockMode interface {
	Encrypt(plain, key string, bc BlockCipher) string
	Decrypt(plain, key string, bc BlockCipher) string
}
