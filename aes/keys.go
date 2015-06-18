package aes

func (k Key128) keyLength() int {
	return 128
}

func (k Key192) keyLength() int {
	return 192
}

func (k Key256) keyLength() int {
	return 256
}
