package aes

func parsePlainText(plain HexString) PlainText {
	var result PlainText
	hexStringToWord(plain, &result)
	return result
}

func toHexString(cipher CipherText) HexString {
	return wordToHexString(cipher)
}

func EncryptHex(key, plain HexString) HexString {
	return toHexString(Encrypt(parseKey(key), parsePlainText(plain)))
}

func stateFrom(plain PlainText) state {
	result := state{}
	copy(result[:], wordsToBytes(plain))
	return transposeState(result)
}

func transposeState(s state) state {
	return state{
		s[0], s[4], s[8], s[12],
		s[1], s[5], s[9], s[13],
		s[2], s[6], s[10], s[14],
		s[3], s[7], s[11], s[15],
	}
}

func stateToCipherText(s state) CipherText {
	r := CipherText{}
	inv := transposeState(s)
	bytesToWord(inv[:], &r)
	return r
}

func Encrypt(key Key, plain PlainText) CipherText {
	schedule := scheduleFor(key)
	state := stateFrom(plain)

	state = addRoundKey(state, schedule.round(0))
	numRounds := key.aesConfiguration().rounds

	for i := 1; i < numRounds; i++ {
		state = addRoundKey(
			mixColumns(
				shiftRows(
					subBytes(state))),
			schedule.round(i))
	}

	state = addRoundKey(shiftRows(subBytes(state)), schedule.round(numRounds))
	return stateToCipherText(state)
}
