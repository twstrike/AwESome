package aes

func EncryptHex(key, plain HexString) HexString {
	return toHexString(Encrypt(parseKey(key), plain.toBlock()))
}

func Encrypt(key Key, plain Block) Block {
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
	return state.toBlock()
}
