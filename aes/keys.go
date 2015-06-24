package aes

func (k Key128) aesConfiguration() AesConfiguration {
	return Aes128
}

func (k Key192) aesConfiguration() AesConfiguration {
	return Aes192
}

func (k Key256) aesConfiguration() AesConfiguration {
	return Aes256
}

func parseKey128(key HexString) Key {
	var result Key128
	hexStringToWord(key, &result)
	return result
}

func parseKey192(key HexString) Key {
	var result Key192
	hexStringToWord(key, &result)
	return result
}

func parseKey256(key HexString) Key {
	var result Key256
	hexStringToWord(key, &result)
	return result
}

func parseKey(key HexString) Key {
	switch len(string(key)) {
	case 32:
		return parseKey128(key)
	case 48:
		return parseKey192(key)
	case 64:
		return parseKey256(key)
	}

	panic("wrong key length")
}
