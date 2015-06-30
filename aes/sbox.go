package aes

import (
	"github.com/twstrike/AwESome/rijndael"
	"github.com/twstrike/AwESome/util"
)

func applySBox(b byte) byte {
	return affineTrans(rijndael.Inv(b))
}

func affineTrans(b byte) byte {
	// See AES spec, formula 5.1
	return b ^ util.RotRight(b, 4) ^ util.RotRight(b, 5) ^ util.RotRight(b, 6) ^ util.RotRight(b, 7) ^ 0x63
}

func applyInvSBox(b byte) byte {
	return rijndael.Inv(invAffineTrans(b))
}

func invAffineTrans(b byte) byte {
	// This is based on the matrix representation of the inverse affine transf
	var A, B, C byte

	A |= util.RotLeft(b, 1) & 0xc1
	A |= (b >> 2) & 0x38
	A |= (b >> 5) & 0x06

	B |= util.RotLeft(b, 3) & 0xc1
	B |= (b << 1) & 0x38
	B |= (b >> 2) & 0x06

	C |= util.RotLeft(b, 6) & 0xc1
	C |= (b << 3) & 0x38
	C |= (b << 1) & 0x06

	return A ^ B ^ C ^ 0x05
}
