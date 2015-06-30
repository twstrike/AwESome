package hash

import (
	"encoding/binary"
	"io"

	"github.com/twstrike/AwESome/util"
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
	return (b & c) | ((^b) & d)
}

func f20to39(b, c, d uint32) uint32 {
	return b ^ c ^ d
}

func f40to59(b, c, d uint32) uint32 {
	return (b & c) | (b & d) | (c & d)
}

func f60to79(b, c, d uint32) uint32 {
	return b ^ c ^ d
}

func fi(i int, b, c, d uint32) uint32 {
	switch {
	case 0 <= i && i <= 19:
		return f0to19(b, c, d)
	case 20 <= i && i <= 39:
		return f20to39(b, c, d)
	case 40 <= i && i <= 59:
		return f40to59(b, c, d)
	case 60 <= i && i <= 79:
		return f60to79(b, c, d)
	}
	return 0
}

const K0to19 = uint32(0x5A827999)
const K20to39 = uint32(0x6ED9EBA1)
const K40to59 = uint32(0x8F1BBCDC)
const K60to79 = uint32(0xCA62C1D6)

func ki(i int) uint32 {
	switch {
	case 0 <= i && i <= 19:
		return K0to19
	case 20 <= i && i <= 39:
		return K20to39
	case 40 <= i && i <= 59:
		return K40to59
	case 60 <= i && i <= 79:
		return K60to79
	}
	return 0
}

// Fills up the buffer from the reader. The only case when it reads less is when EOF is encountered
func readExactly(into []byte, r io.Reader) (read int, err error) {
	for read < len(into) && err == nil {
		var n int
		n, err = r.Read(into[read:])
		read += n
	}
	return
}

func (ctx *sha1Context) init() {
	ctx.H0 = 0x67452301
	ctx.H1 = 0xEFCDAB89
	ctx.H2 = 0x98BADCFE
	ctx.H3 = 0x10325476
	ctx.H4 = 0xC3D2E1F0
}

func (ctx *sha1Context) update(mN [sha1BlockSizeInBytes]byte) {
	// runs the core of the algorithm
	for t := 0; t < 16; t++ {
		ctx.W[t] = util.BytesToUint32([4]byte{mN[t*4], mN[t*4+1], mN[t*4+2], mN[t*4+3]})
	}
	for t := 16; t < 80; t++ {
		ctx.W[t] = util.RotLeftUint32(ctx.W[t-3]^ctx.W[t-8]^ctx.W[t-14]^ctx.W[t-16], 1)
	}

	ctx.A, ctx.B, ctx.C, ctx.D, ctx.E = ctx.H0, ctx.H1, ctx.H2, ctx.H3, ctx.H4

	for t := 0; t < 80; t++ {
		ctx.temp = sumUint32Modulo([]uint32{util.RotLeftUint32(ctx.A, 5), fi(t, ctx.B, ctx.C, ctx.D), ctx.E, ki(t), ctx.W[t]})
		ctx.E = ctx.D
		ctx.D = ctx.C
		ctx.C = util.RotLeftUint32(ctx.B, 30)
		ctx.B = ctx.A
		ctx.A = ctx.temp
	}

	ctx.H0 = addUint32Modulo(ctx.H0, ctx.A)
	ctx.H1 = addUint32Modulo(ctx.H1, ctx.B)
	ctx.H2 = addUint32Modulo(ctx.H2, ctx.C)
	ctx.H3 = addUint32Modulo(ctx.H3, ctx.D)
	ctx.H4 = addUint32Modulo(ctx.H4, ctx.E)
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
	atEnd = err == io.EOF

	switch {
	case l == 0:
		if err != io.EOF {
			panic("Shouldn't happen")
		}

		if sha1.needPadding {
			sha1.addPadding(buffer)
			hasContent = true
		}
		hasContent = false
	case l == sha1BlockSizeInBytes:
		sha1.needPadding = false
		hasContent = true
	case l < (sha1BlockSizeInBytes - 9):
		if err != io.EOF {
			panic("Shouldn't happen")
		}

		buffer[l] = 0x80
		sha1.addPadding(buffer)
		hasContent = true
	case l < sha1BlockSizeInBytes:
		if err != io.EOF {
			panic("Shouldn't happen")
		}

		buffer[l] = 0x80
		sha1.needPadding = true
		hasContent, atEnd = true, false
	default:
		panic("Shouldn't happen")
	}

	return
}

func (sha1 *sha1Reader) sum() [sha1OutputSizeInBytes]byte {
	ctx := sha1Context{}
	ctx.init()
	into := [sha1OutputSizeInBytes]byte{}
	readExactly(into[:], sha1.reader)

	return [sha1OutputSizeInBytes]byte{}
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
