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

func stateFrom(block Block) state {
	result := state{}
	copy(result[:], wordsToBytes(block))
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
	return CipherText(stateToBlockText(s))
}

func stateToPlainText(s state) PlainText {
	return PlainText(stateToBlockText(s))
}

func stateToBlockText(s state) Block {
	r := Block{}
	inv := transposeState(s)
	bytesToWord(inv[:], &r)
	return r
}

func Encrypt(key Key, plain PlainText) CipherText {
	schedule := scheduleFor(key)
	state := stateFrom(Block(plain))

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

func Decrypt(key Key, cipher CipherText) PlainText {
	schedule := scheduleFor(key)
	state := stateFrom(Block(cipher))

	numRounds := key.aesConfiguration().rounds
	state = addRoundKey(state, schedule.round(numRounds))

	for i := numRounds - 1; i > 0; i-- {
		state = invMixColumns(
			addRoundKey(
				invSubBytes(
					invShiftRows(state)),
				schedule.round(i)))
	}

	state = addRoundKey(invSubBytes(invShiftRows(state)), schedule.round(0))
	return stateToPlainText(state)
}
