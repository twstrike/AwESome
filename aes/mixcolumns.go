package aes

import (
	"github.com/twstrike/AwESome/rijndael"
)

func mul1(b byte) byte {
	return b
}

func mul2(b byte) byte {
	h := byte(int8(b) >> 7)
	return (b << 1) ^ (0x1B & h)
}

func mul3(b byte) byte {
	return b ^ mul2(b)
}

func mixOneColumn(s state, col int) state {
	a0 := s[0+col]
	a1 := s[4+col]
	a2 := s[8+col]
	a3 := s[12+col]

	result := s
	result[0+col] = mul2(a0) ^ mul3(a1) ^ mul1(a2) ^ mul1(a3)
	result[4+col] = mul1(a0) ^ mul2(a1) ^ mul3(a2) ^ mul1(a3)
	result[8+col] = mul1(a0) ^ mul1(a1) ^ mul2(a2) ^ mul3(a3)
	result[12+col] = mul3(a0) ^ mul1(a1) ^ mul1(a2) ^ mul2(a3)
	return result
}

func invMixOneColumn(s state, col int) state {
	a0 := s[0+col]
	a1 := s[4+col]
	a2 := s[8+col]
	a3 := s[12+col]

	result := s
	result[0+col] = rijndael.Mul(0x0e, a0) ^ rijndael.Mul(0x0b, a1) ^ rijndael.Mul(0x0d, a2) ^ rijndael.Mul(0x09, a3)
	result[4+col] = rijndael.Mul(0x09, a0) ^ rijndael.Mul(0x0e, a1) ^ rijndael.Mul(0x0b, a2) ^ rijndael.Mul(0x0d, a3)
	result[8+col] = rijndael.Mul(0x0d, a0) ^ rijndael.Mul(0x09, a1) ^ rijndael.Mul(0x0e, a2) ^ rijndael.Mul(0x0b, a3)
	result[12+col] = rijndael.Mul(0x0b, a0) ^ rijndael.Mul(0x0d, a1) ^ rijndael.Mul(0x09, a2) ^ rijndael.Mul(0x0e, a3)
	return result
}

func mixColumns(s state) state {
	result := s
	for i := 0; i < 4; i++ {
		result = mixOneColumn(result, i)
	}
	return result
}

func invMixColumns(s state) state {
	result := s
	for i := 0; i < 4; i++ {
		result = invMixOneColumn(result, i)
	}
	return result
}
