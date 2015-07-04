package aes

// DecryptBytes decrypt the bytes in the cipher text with the key and returns the result
func DecryptBytes(key, cipher []byte) []byte {
	k := parseKeyFromBytes(key)
	c := Block{}
	bytesToWord(cipher, &c)

	return wordsToBytes(Decrypt(k, c))
}

// Decrypt will decrypt the bytes in the cipher text with the key and returns the result
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
