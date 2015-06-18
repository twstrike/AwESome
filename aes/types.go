package aes

type HexString string

type word uint32
type Key128 [4]word
type Key192 [6]word
type Key256 [8]word

type PlainText [4]word
type CipherText [4]word

type state [4 * 4]byte

type roundSchedule [4]word
type keySchedule128 [Nr128 + 1]roundSchedule
