package util

func RotRight(b, n byte) byte {
	n = n % 8
	return (b >> n) | (b << (8 - n))
}

func RotLeft(b, n byte) byte {
	n = n % 8
	return (b << n) | (b >> (8 - n))
}

func RotRightUint32(b, n uint32) uint32 {
	n = n % 32
	return (b >> n) | (b << (32 - n))
}

func RotLeftUint32(b, n uint32) uint32 {
	n = n % 32
	return (b << n) | (b >> (32 - n))
}

func BytesToUint32(b [4]byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func AddUint32Modulo(a, b uint32) uint32 {
	return uint32((uint64(a) + uint64(b)) % uint64(0x0000000100000000))
}
