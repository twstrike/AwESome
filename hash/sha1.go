package hash

import (
	"encoding/binary"
	"io"
)

type SHA1 struct{}
type sha1Reader struct {
	reader      io.Reader
	currentSize uint64
	needPadding bool
}

type sha1Context struct {
	A, B, C, D, E      uint32
	H0, H1, H2, H3, H4 uint32
	W                  [80]uint32
	temp               uint32
}

func (sha1 SHA1) Sum(r io.Reader) []byte {
	reader := sha1Reader{r, 0, false}
	result := reader.sum()
	return result[:]
}

const sha1OutputSize = 160
const sha1OutputSizeInBytes = sha1OutputSize / 8
const sha1BlockSize = 512
const sha1BlockSizeInBytes = sha1BlockSize / 8

type sha1Block [sha1BlockSizeInBytes]byte

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
	for read < len(into) && err == nil {
		var n int
		n, err = r.Read(into[read:])
		read += n
	}
	return
}

func (ctx *sha1Context) update(mN [sha1BlockSizeInBytes]byte) {
	// runs the core of the algorithm
}

func (ctx *sha1Context) final() [sha1OutputSizeInBytes]byte {
	return [sha1OutputSizeInBytes]byte{}
}

func (sha1 *sha1Reader) addPadding(buffer *sha1Block) {
	w := NewSliceWriter(buffer[len(buffer)-8:])
	binary.Write(w, binary.BigEndian, sha1.currentSize)
}

func (sha1 *sha1Reader) readWithPadding(buffer *sha1Block) (hasContent bool, atEnd bool) {
	l, err := readExactly(buffer[:], sha1.reader)
	sha1.currentSize += uint64(l * 8)

	switch {
	case l == 0:
		if err == io.EOF {
			if sha1.needPadding {
				sha1.addPadding(buffer)
				return true, true
			}
			return false, true
		}

		panic("Shouldn't happen")
	case l == sha1BlockSizeInBytes:
		sha1.needPadding = false
		// add content
		return true, err == io.EOF
	case l < (sha1BlockSizeInBytes - 9):
		if err == io.EOF {
			buffer[l] = 0x80
			sha1.addPadding(buffer)
			return true, true
		}
		panic("Shouldn't happen")
	case l < sha1BlockSizeInBytes:
		if err == io.EOF {
			buffer[l] = 0x80
			sha1.needPadding = true
			return true, false
		}
		panic("Shouldn't happen")
	}

	panic("Shouldn't happen")
}

func (sha1 *sha1Reader) sum() [sha1OutputSizeInBytes]byte {
	return [sha1OutputSizeInBytes]byte{}
}
