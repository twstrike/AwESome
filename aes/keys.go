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

func (key Key128) newKeySchedule() keySchedule {
	result := make(KeySchedule, key.aesConfiguration().rounds+1)
	s := newScheduleFor(key[:], key.aesConfiguration().rounds)

	for i := 0; i < key.aesConfiguration().rounds+1; i++ {
		result[i] = roundSchedule{s[i*4+0], s[i*4+1], s[i*4+2], s[i*4+3]}
	}

	return result
}

func (key Key192) newKeySchedule() keySchedule {
	result := make(KeySchedule, key.aesConfiguration().rounds+1)
	s := newScheduleFor(key[:], key.aesConfiguration().rounds)

	for i := 0; i < key.aesConfiguration().rounds+1; i++ {
		result[i] = roundSchedule{s[i*4+0], s[i*4+1], s[i*4+2], s[i*4+3]}
	}

	return result
}

func (key Key256) newKeySchedule() keySchedule {
	result := make(KeySchedule, key.aesConfiguration().rounds+1)
	s := newScheduleFor(key[:], key.aesConfiguration().rounds)

	for i := 0; i < key.aesConfiguration().rounds+1; i++ {
		result[i] = roundSchedule{s[i*4+0], s[i*4+1], s[i*4+2], s[i*4+3]}
	}

	return result
}

func (s KeySchedule) round(i int) roundSchedule {
	return s[i]
}
