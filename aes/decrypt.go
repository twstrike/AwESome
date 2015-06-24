package aes

func DecryptHex(key, cipher HexString) HexString {
	return Decrypt(parseKey(key), cipher.toBlock()).toHexString()
}

func Decrypt(key Key, cipher Block) Block {
	schedule := scheduleFor(key)
	state := cipher.toState()

	numRounds := key.aesConfiguration().rounds
	state = addRoundKey(state, schedule[numRounds])

	for i := numRounds - 1; i > 0; i-- {
		state = invMixColumns(
			addRoundKey(
				invSubBytes(
					invShiftRows(state)),
				schedule[i]))
	}

	state = addRoundKey(invSubBytes(invShiftRows(state)), schedule[0])
	return state.toBlock()
}
