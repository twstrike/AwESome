package aes

func EncryptHex(key, plain HexString) HexString {
	return Encrypt(parseKey(key), plain.toBlock()).toHexString()
}

func Encrypt(key Key, plain Block) Block {
	schedule := scheduleFor(key)
	state := plain.toState()

	state = addRoundKey(state, schedule[0])
	numRounds := key.aesConfiguration().rounds

	for i := 1; i < numRounds; i++ {
		state = addRoundKey(
			mixColumns(
				shiftRows(
					subBytes(state))),
			schedule[i])
	}

	state = addRoundKey(shiftRows(subBytes(state)), schedule[numRounds])
	return state.toBlock()
}
