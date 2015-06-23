package aes

type HexString string

type word uint32
type Key128 [4]word
type Key192 [6]word
type Key256 [8]word

type Block [4]word

type state [4 * 4]byte

type roundSchedule [4]word
type keySchedule128 [Nr128 + 1]roundSchedule

type AesConfiguration struct {
	keyLength int
	rounds    int
}

type keySchedule interface {
	round(i int) roundSchedule
}

type Key interface {
	aesConfiguration() AesConfiguration
	newKeySchedule() keySchedule
}
