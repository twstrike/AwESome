package hash

import "io"

type SHA1 struct{}
type sha1Reader struct {
	reader      io.Reader
	currentSize uint64
}

func (sha1 SHA1) Sum(r io.Reader) []byte {
	reader := sha1Reader{r, 0}
	return reader.sum()
}

const sha1BlockSize = 512
const sha1BlockSizeInBytes = sha1BlockSize / 8

func circularLeftShift(w uint32, n int) uint32 {
	return w
}

func f0to19(b, c, d uint32) uint32 {
	return 0
}

func f20to39(b, c, d uint32) uint32 {
	return 0
}

func f40to59(b, c, d uint32) uint32 {
	return 0
}

func f60to79(b, c, d uint32) uint32 {
	return 0
}

const K0to19 = uint32(0x5A827999)
const K20to39 = uint32(0x6ED9EBA1)
const K40to59 = uint32(0x8F1BBCDC)
const K60to79 = uint32(0xCA62C1D6)

// Fills up the buffer from the reader. The only case when it reads less is when EOF is encountered
func readExactly(into []byte, r io.Reader) (read int, err error) {
	return 0, nil
}

func (sha1 *sha1Reader) readWithPadding(buffer [sha1BlockSizeInBytes]byte) (atEnd bool) {
	return true
}

func (sha1 *sha1Reader) sum() []byte {
	return []byte{}
}
