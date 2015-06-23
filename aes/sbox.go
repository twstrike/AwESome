package aes

import (
	"github.com/twstrike/AwESome/rijndael"
)

func applySBox(b byte) byte {
	return affineTrans(rijndael.Inv(b))
}

func affineTrans(b byte) byte {
	// See AES spec, formula 5.1
	return b ^ rotRight(b, 4) ^ rotRight(b, 5) ^ rotRight(b, 6) ^ rotRight(b, 7) ^ 0x63
}

func rotRight(b, n byte) byte {
	n = n % 8
	return (b >> n) | (b << (8 - n))
}

func applyInvSBox(b byte) byte {
	return rijndael.Inv(invAffineTrans(b))
}

func invAffineTrans(b byte) byte {
	// This is based on the matrix representation of the inverse affine transf
	var A, B, C byte

	A |= rotLeft(b, 1) & 0xc1
	A |= (b >> 2) & 0x38
	A |= (b >> 5) & 0x06

	B |= rotLeft(b, 3) & 0xc1
	B |= (b << 1) & 0x38
	B |= (b >> 2) & 0x06

	C |= rotLeft(b, 6) & 0xc1
	C |= (b << 3) & 0x38
	C |= (b << 1) & 0x06

	return A ^ B ^ C ^ 0x05
}

func rotLeft(b, n byte) byte {
	n = n % 8
	return (b << n) | (b >> (8 - n))
}
