package aes

func subBytes(s state) (out state) {
	for i, b := range s {
		out[i] = applySBox(b)
	}
	return
}

func invSubBytes(s state) (out state) {
	for i, b := range s {
		out[i] = applyInvSBox(b)
	}
	return
}
