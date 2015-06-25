package block

type BlockCipher interface {
	BlockSize() int
	Encrypt(block, key []byte) []byte
	Decrypt(block, key []byte) []byte
}
