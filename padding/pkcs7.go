package padding

func PKCS7(block []byte, size int) []byte {
	pads := size - len(block)%size
	padded := make([]byte, len(block)+pads)
	copy(padded[:], block)

	for i := len(block); i < len(padded); i++ {
		padded[i] = byte(pads)
	}
	return padded
}
