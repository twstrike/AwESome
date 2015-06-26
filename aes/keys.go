package aes

import (
	"fmt"
)

func (k Key128) aesConfiguration() aesConfiguration {
	return aes128
}

func (k Key192) aesConfiguration() aesConfiguration {
	return aes192
}

func (k Key256) aesConfiguration() aesConfiguration {
	return aes256
}

func parseKey128(key []byte) Key {
	var result Key128
	bytesToWord(key, &result)
	return result
}

func parseKey192(key []byte) Key {
	var result Key192
	bytesToWord(key, &result)
	return result
}

func parseKey256(key []byte) Key {
	var result Key256
	bytesToWord(key, &result)
	return result
}

func parseKeyFromBytes(key []byte) Key {
	switch len(key) {
	case 16:
		return parseKey128(key)
	case 24:
		return parseKey192(key)
	case 32:
		return parseKey256(key)
	}

	panic(fmt.Sprintf("invalid key length: %d", len(key)))
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
