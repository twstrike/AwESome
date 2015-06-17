package aes

import(
    "strings"
)

type HexString string

type word uint32
type Key128 [4]word
type Key192 [6]word
type Key256 [8]word

type PlainText [4]word
type CipherText [4]word

type state [4*4]byte

func parseKey(key HexString) Key128 {
	var result Key128
    noSpacesStr := strings.Replace(string(key), " ", "", -1)
    hexStringToWord(HexString(noSpacesStr), &result)
	return result
}

func parsePlainText(plain HexString) PlainText {
	// TODO: implement
	return PlainText{0, 0, 0, 0}
}

func toHexString(cipher CipherText) HexString {
	// TODO: implement
	return HexString("placeholder")
}

func EncryptHex(key, plain HexString) HexString {
	return toHexString(Encrypt128(parseKey(key), parsePlainText(plain)))
}

func stateFrom(plain PlainText) state {
	// TODO: implement
     return state{}
}

func stateToCipherText(s state) CipherText {
	// TODO: implement
     return CipherText{}
}

func Encrypt128(key Key128, plain PlainText) CipherText {
	schedule := scheduleFor(key)
	state := stateFrom(plain)

	state = addRoundKey(state, schedule[0])
	for i := 1; i < Nr128; i++ {
	    state = addRoundKey(
	    	  mixColumns(
			shiftRows(
				subBytes(state))),
		schedule[i])
	}

	state = addRoundKey(shiftRows(subBytes(state)), schedule[Nr128])
	return stateToCipherText(state)
}
