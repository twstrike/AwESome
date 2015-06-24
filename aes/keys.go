package aes

func (k Key128) aesConfiguration() aesConfiguration {
	return aes128
}

func (k Key192) aesConfiguration() aesConfiguration {
	return aes192
}

func (k Key256) aesConfiguration() aesConfiguration {
	return aes256
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

func (k Key128) newKeySchedule() keySchedule {
	s := keyExpand(k[:], k.aesConfiguration().rounds)
	return collectRoundSchedule(s, k)
}

func (k Key192) newKeySchedule() keySchedule {
	s := keyExpand(k[:], k.aesConfiguration().rounds)
	return collectRoundSchedule(s, k)
}

func (k Key256) newKeySchedule() keySchedule {
	s := keyExpand(k[:], k.aesConfiguration().rounds)
	return collectRoundSchedule(s, k)
}
