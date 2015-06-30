package block

type Cipher interface {
	BlockSize() int
	Encrypt(key, plain []byte) []byte
	Decrypt(key, cipher []byte) []byte
}
