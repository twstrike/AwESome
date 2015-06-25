package block

type BlockCipher interface {
	BlockSize() int
	Encrypt(key, plain []byte) []byte
	Decrypt(key, cipher []byte) []byte
}
