package aes

type HexString string

type word uint32
type Key128 [4]word
type Key192 [6]word
type Key256 [8]word

type Block [4]word

type state [4 * 4]byte

type roundSchedule [4]word
type keySchedule []roundSchedule

type aesConfiguration struct {
	keyLength int
	rounds    int
}

type Key interface {
	aesConfiguration() aesConfiguration
	newKeySchedule() keySchedule
}
