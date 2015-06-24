package aes

import (
	"github.com/twstrike/AwESome/rijndael"
)

func newScheduleFor(key []word, Nr uint) []word {
	Nb := uint(4)
	Nk := len(key)
	result := make([]word, Nb*(Nr+1))

	copy(result[:], key)

	for i := Nk; i < len(result); i++ {
		temp := result[i-1]
		if i%Nk == 0 {
			temp = subWord(rotWord(temp)) ^ rcon(i/Nk)
		} else if Nk > 6 && i%Nk == 4 {
			temp = subWord(temp)
		}

		result[i] = result[i-Nk] ^ temp
	}

	return result
}

func (key Key128) newKeySchedule() keySchedule {
	result := keySchedule128{}
	s := newScheduleFor(key[:], Nr128)

	for i := 0; i < len(result); i++ {
		result[i] = roundSchedule{s[i*4+0], s[i*4+1], s[i*4+2], s[i*4+3]}
	}

	return result
}

func (s keySchedule128) round(i int) roundSchedule {
	return s[i]
}

func subWord(w word) word {
	out := word(0)
	out |= word(applySBox(byte(w>>24))) << 24
	out |= word(applySBox(byte(w>>16))) << 16
	out |= word(applySBox(byte(w>>8))) << 8
	out |= word(applySBox(byte(w>>0))) << 0
	return out
}

func rotWord(w word) word {
	return w<<8 | (w >> 24)
}

func rcon(i int) word {
	return word(rijndael.Exp(2, byte(i-1))) << 24
}

func scheduleFor(key Key) keySchedule {
	return key.newKeySchedule()
}
