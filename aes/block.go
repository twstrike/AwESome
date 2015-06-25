package aes

func (b Block) toState() state {
	result := state{}
	copy(result[:], wordsToBytes(b))
	return result.transpose()
}
