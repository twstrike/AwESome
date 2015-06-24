package block

type ECB struct{}

type BlockMode interface {
	Encrypt(plain, key string) string
	Decrypt(plain, key string) string
}
