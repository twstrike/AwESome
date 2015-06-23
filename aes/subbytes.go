package aes

func subBytes(s state) state {
	out := state{}

	for i, b := range s {
		out[i] = applySBox(b)
	}

	return out
}

func invSubBytes(s state) state {
	out := state{}

	for i, b := range s {
		out[i] = applyInvSBox(b)
	}

	return out
}
