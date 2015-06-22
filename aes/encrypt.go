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
	return result
}

func stateToCipherText(s state) CipherText {
	result := CipherText{}
	bytesToWord(s[:], &result)
	return result
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
