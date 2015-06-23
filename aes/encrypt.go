package aes

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
	return state.toCipherText()
}
