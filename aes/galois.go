package aes

func galoisAdd(i, j byte) byte {
	return i ^ j
}

func galoisSub(i, j byte) byte {
	return i ^ j
}

func galoisMul(i, j byte) byte {
	return modulo(multiplication(i, j), 0x11B)
}

func galoisExp(x, n byte) byte {
	result := byte(1)

	for i := byte(0); i < n; i++ {
		result = galoisMul(result, x)
	}

	return result
}

func multiplication(i, j byte) uint16 {
	var out uint16 = 0
	a := uint16(i)
	b := uint16(j)

	for place := uint16(0); place < 8; place++ {
		var k uint16 = 1 << place
		// if bit K in b is set, mask will be 0xFF, otherwise 0x00 - this avoids branching
		mask := uint16((int16(k&b) << (15 - place)) >> 15)
		out = out ^ ((a << place) & mask)
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
