package util

// RotRight rotates b right by n
func RotRight(b, n byte) byte {
	n = n % 8
	return (b >> n) | (b << (8 - n))
}

// RotLeft rotates b left by n
func RotLeft(b, n byte) byte {
	n = n % 8
	return (b << n) | (b >> (8 - n))
}

// RotRightUint32 rotates b right by n
func RotRightUint32(b, n uint32) uint32 {
	n = n % 32
	return (b >> n) | (b << (32 - n))
}

// RotLeftUint32 rotates b left by n
func RotLeftUint32(b, n uint32) uint32 {
	n = n % 32
	return (b << n) | (b >> (32 - n))
}

// BytesToUint32 turns four bytes b into a uint32
func BytesToUint32(b [4]byte) uint32 {
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

// AddUint32Modulo adds two uint32 together
func AddUint32Modulo(a, b uint32) uint32 {
	return uint32((uint64(a) + uint64(b)) % uint64(0x0000000100000000))
}
