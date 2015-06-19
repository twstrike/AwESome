package aes

import (
  "github.com/twstrike/AwESome/rijndael"
)

func (key Key128) newKeySchedule() keySchedule {
	result := keySchedule128{}
	result[0][0] = key[0]
	result[0][1] = key[1]
	result[0][2] = key[2]
	result[0][3] = key[3]

	for i := 1; i <= Nr128; i++ {
		prev := result[i-1]
		w0 := prev[0] ^ (subWord(rotWord(prev[3])) ^ word(rcon(i)))
		w1 := prev[1] ^ w0
		w2 := prev[2] ^ w1
		w3 := prev[3] ^ w2
		result[i][0] = w0
		result[i][1] = w1
		result[i][2] = w2
		result[i][3] = w3
	}

	return &result
}

func (s keySchedule128) round(i int) roundSchedule {
	return s[i]
}

func subWord(w word) word {
	return w
}

func rotWord(w word) word {
	return w<<8 | (w >> 24)
}

func rcon(i int) byte {
	return rijndael.Exp(2, byte(i-1))
}
func scheduleFor(key Key) keySchedule {
	return key.newKeySchedule()
}
