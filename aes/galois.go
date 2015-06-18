package aes

func galoisMul(i, j byte) byte {
	return modulo(multiplication(i, j), 0x11B)
}

func galoisExp(x, n byte) byte {
	result := byte(1)

	for i:=byte(0); i<n; i++ {
		result = galoisMul(result, x)
	}

	return result
}
