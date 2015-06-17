package aes

func leftShift(s *state, at, n int) {
	left := make([]byte, n)
	copy(left, s[at:])
	copy(s[at:(at+(4-n))], s[(at+n):(at+n+(4-n))])
	copy(s[(at+(4-n)):(at+n+(4-n))], left)
}

func shiftRows(s state) state {
	result := s
	leftShift(&result, 1*4, 1)
	leftShift(&result, 2*4, 2)
	leftShift(&result, 3*4, 3)
	return result
}
