package hash

import (
	"fmt"

	"github.com/twstrike/AwESome/util"
)

const k0to19 = uint32(0x5A827999)
const k20to39 = uint32(0x6ED9EBA1)
const k40to59 = uint32(0x8F1BBCDC)
const k60to79 = uint32(0xCA62C1D6)

type sha1Context struct {
	A, B, C, D, E      uint32
	H0, H1, H2, H3, H4 uint32
	W                  [80]uint32
	temp               uint32
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
	return uint32ToSHA1Output([5]uint32{ctx.H0, ctx.H1, ctx.H2, ctx.H3, ctx.H4})
}

func ki(i int) uint32 {
	switch {
	default:
		panic(fmt.Sprintf("invalid i: ", i))
	case 0 <= i && i <= 19:
		return k0to19
	case 20 <= i && i <= 39:
		return k20to39
	case 40 <= i && i <= 59:
		return k40to59
	case 60 <= i && i <= 79:
		return k60to79
	}
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
	default:
		panic(fmt.Sprintf("invalid i: ", i))
	case 0 <= i && i <= 19:
		return f0to19(b, c, d)
	case 20 <= i && i <= 39:
		return f20to39(b, c, d)
	case 40 <= i && i <= 59:
		return f40to59(b, c, d)
	case 60 <= i && i <= 79:
		return f60to79(b, c, d)
	}
}
