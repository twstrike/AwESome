package aes

import (
	"encoding/hex"
)

func (b Block) toState() state {
	result := state{}
	copy(result[:], wordsToBytes(b))
	return result.transpose()
}

func (b Block) toHexString() HexString {

  	encoded := hex.EncodeToString(wordsToBytes(b))
  	return HexString(encoded)
}
