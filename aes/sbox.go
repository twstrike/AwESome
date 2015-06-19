package aes

import (
	"github.com/twstrike/AwESome/rijndael"
)

func applySBox(b byte) byte {
	return affineTrans(rijndael.Inv(b))
}

func affineTrans(b byte) byte {
	return b ^ rotRight(b, 4) ^ rotRight(b, 5) ^ rotRight(b, 6) ^ rotRight(b, 7) ^ 0x63
}

func rotRight(b, n byte) byte {
	n = n % 8
	return (b >> n) | (b << (8 - n))
}
