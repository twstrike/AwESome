package aes

type word uint32

// Key128 represents a 128 bit key
type Key128 [4]word

// Key192 represents a 192 bit key
type Key192 [6]word

// Key256 represents a 256 bit key
type Key256 [8]word

// Block represents a block
type Block [4]word

type state [4 * 4]byte

type roundSchedule [4]word
type keySchedule []roundSchedule

type aesConfiguration struct {
	keyLength int
	rounds    int
}

// Key contains configuration information for a specific AES instantiation
type Key interface {
	aesConfiguration() aesConfiguration
	newKeySchedule() keySchedule
}
