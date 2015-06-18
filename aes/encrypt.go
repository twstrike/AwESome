package aes

func parseKey(key HexString) Key128 {
	var result Key128
	hexStringToWord(key, &result)
	return result
}

func parsePlainText(plain HexString) PlainText {
	var result PlainText
	hexStringToWord(plain, &result)
	return result
}

func toHexString(cipher CipherText) HexString {
	return wordToHexString(cipher)
}

func EncryptHex(key, plain HexString) HexString {
	return toHexString(Encrypt128(parseKey(key), parsePlainText(plain)))
}

func stateFrom(plain PlainText) state {
	// TODO: implement
	return state{}
}

func stateToCipherText(s state) CipherText {
	// TODO: implement
	return CipherText{}
}

func Encrypt128(key Key128, plain PlainText) CipherText {
	schedule := scheduleFor(key)
	state := stateFrom(plain)

	state = addRoundKey(state, schedule[0])
	for i := 1; i < Nr128; i++ {
		state = addRoundKey(
			mixColumns(
				shiftRows(
					subBytes(state))),
			schedule[i])
	}

	state = addRoundKey(shiftRows(subBytes(state)), schedule[Nr128])
	return stateToCipherText(state)
}
