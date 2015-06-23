package aes

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
	return state.toPlainText()
}
