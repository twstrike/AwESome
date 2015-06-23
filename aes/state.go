package aes

func (s state) transpose() state {
	return state{
		s[0], s[4], s[8], s[12],
		s[1], s[5], s[9], s[13],
		s[2], s[6], s[10], s[14],
		s[3], s[7], s[11], s[15],
	}
}

func (s state) toCipherText() CipherText {
	return CipherText(s.toBlock())
}

func (s state) toPlainText() PlainText {
	return PlainText(s.toBlock())
}

func (s state) toBlock() Block {
	r := Block{}
	inv := s.transpose()
	bytesToWord(inv[:], &r)
	return r
}
