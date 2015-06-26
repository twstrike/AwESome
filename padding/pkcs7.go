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

func UndoPKCS7(padded []byte) []byte {
	pads := int(padded[len(padded)-1])
	unpadded := make([]byte, len(padded)-pads)
	copy(unpadded, padded[:len(padded)-pads])
	return unpadded
}
