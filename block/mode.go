package block

type Mode interface {
	Encrypt(plain, key []byte, bc Cipher) []byte
	Decrypt(plain, key []byte, bc Cipher) []byte
}
