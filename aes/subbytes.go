package aes

func subBytes(s state) state {
	var newState state
	for i, stateline := range s {
		for j, b := range stateline {
			newState[i][j] = affineTrans(b)
		}
	}
	return newState
}

func affineTrans(b byte) byte {
	const affin = 0x8F

	return 0x11
}

func helper(place, a, out, b uint16) uint16 {
	var k uint16 = 1 << place
	// if bit K in b is set, mask will be 0xFF, otherwise 0x00 - this avoids branching
	mask := uint16((int16(k&b) << (15 - place)) >> 15)
	return out ^ ((a << place) & mask)
}

func multiplication(i, j byte) uint16 {
	var out uint16 = 0
	a := uint16(i)
	b := uint16(j)

	for position := 0; position < 8; i++ {
		out = helper(position, a, out, b)
	}

	return out
}

func modulo(multiplied, poly uint16) byte {
	out := uint16(multiplied)
	b := uint16(poly)
	for nbits(out) >= nbits(b) {
		out = out ^ (b << (nbits(out) - nbits(b)))
	}
	return byte(out)
}

func nbits(a uint16) uint {
	n := uint(0)

	for ; a > 0; a >>= 1 {
		n++
	}

	return n
}
