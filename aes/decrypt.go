package aes

func DecryptHex(key, cipher HexString) HexString {
	return toHexString(Decrypt(parseKey(key), cipher.toBlock()))
}

func Decrypt(key Key, cipher Block) Block {
	schedule := scheduleFor(key)
	state := stateFrom(cipher)

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
	return state.toBlock()
}
