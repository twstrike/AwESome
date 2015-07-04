package block

import "bytes"

// CTR contains the information necessary to run the CTR block mode
type CTR struct {
	IV []byte
}

func incAt(d []byte, i int) {
	d[i]++
	if d[i] == 0 && i > 0 {
		incAt(d, i-1)
	}
}

func inc(d []byte) {
	incAt(d, len(d)-1)
}

// Encrypt encrypts the plain text with the key using the provided block cipher and CTR block mode
func (bm CTR) Encrypt(key, plain []byte, blockCipher Cipher) []byte {
	reader := bytes.NewBuffer(plain)
	cipher := make([]byte, len(plain))
	blockSize := blockCipher.BlockSize() / 8
	counter := make([]byte, blockSize)
	copy(counter[:], bm.IV)

	for i := 0; reader.Len() > 0; i += blockSize {
		stream := blockCipher.Encrypt(key, counter)
		copy(cipher[i:], xor(stream, reader.Next(blockSize)))
		inc(counter)
	}

	return cipher
}

// Decrypt decrypts the cipher text with the key using the provided block cipher and CTR block mode
func (bm CTR) Decrypt(key, cipher []byte, blockCipher Cipher) []byte {
	return bm.Encrypt(key, cipher, blockCipher)
}
