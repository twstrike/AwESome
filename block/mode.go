package block

// Mode represents the operations that can be done on a block mode
type Mode interface {
	Encrypt(plain, key []byte, bc Cipher) []byte
	Decrypt(plain, key []byte, bc Cipher) []byte
}
