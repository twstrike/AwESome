package aes

// EncryptBytes encrypts the bytes in the plain text with the key and returns the result
func EncryptBytes(key, plain []byte) []byte {
	k := parseKeyFromBytes(key)
	p := Block{}
	bytesToWord(plain, &p)

	return wordsToBytes(Encrypt(k, p))
}

// Encrypt will encrypt the bytes in the plain text with the key and returns the result
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
