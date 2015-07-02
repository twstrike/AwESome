package hash

import (
	"encoding/binary"
	"io"
)

const sha1OutputSize = 160
const sha1OutputSizeInBytes = sha1OutputSize / 8
const sha1BlockSize = 512
const sha1BlockSizeInBytes = sha1BlockSize / 8

type sha1Block [sha1BlockSizeInBytes]byte

// SHA1 is a hash.Hash that computes the SHA1 message digest.
type SHA1 struct{}

// Sum returns the SHA1 digest for the message in r.
func (sha1 SHA1) Sum(r io.Reader) []byte {
	message := NewSha1MessageReader(r)
	digest := compute(message)
	return digest[:]
}

func (sha1 *SHA1MessageReader) readWithPadding(buffer *sha1Block) (atEnd bool) {
	_, err := sha1.Read(buffer[:])
	return err == io.EOF
}

func compute(msg *SHA1MessageReader) [sha1OutputSizeInBytes]byte {
	ctx := sha1Context{}
	ctx.init()

	block := sha1Block{}

	for {
		_, err := msg.Read(block[:])
		ctx.update(block)
		if err != nil {
			break
		}
	}
	return ctx.final()
}

func addUint32Modulo(a, b uint32) uint32 {
	return uint32((uint64(a) + uint64(b)) % uint64(0x0000000100000000))
}

func sumUint32Modulo(arr []uint32) uint32 {
	var result = uint32(0)
	for i := range arr {
		result = addUint32Modulo(result, arr[i])
	}
	return result
}

func uint32ToSHA1Output(h [5]uint32) (result [sha1OutputSizeInBytes]byte) {
	w := NewSliceWriter(result[:])
	for i := range h {
		binary.Write(w, binary.BigEndian, h[i])
	}
	return
}
