package hash

import (
	"encoding/binary"
	"io"

	aio "github.com/twstrike/AwESome/io"
)

// SHA1MessageReader implements SHA-1 padding for an io.Reader object.
type SHA1MessageReader struct {
	r    *aio.TrustworthyReader
	size uint64
}

// NewSha1MessageReader returns new SHA1MessageReader wrapping r
func NewSha1MessageReader(r io.Reader) *SHA1MessageReader {
	return &SHA1MessageReader{r: aio.NewTrustworthyReader(r)}
}

func (r *SHA1MessageReader) addMessageLength(buffer []byte) {
	w := NewSliceWriter(buffer[len(buffer)-8:])
	binary.Write(w, binary.BigEndian, r.size)
}

func emptySHA1Block(p []byte) {
	for i := range p {
		p[i] = 0
	}
}

// Read reads exactly sha1BlockSizeInBytes bytes into p
func (r *SHA1MessageReader) Read(p []byte) (int, error) {
	if len(p) < sha1BlockSizeInBytes {
		return 0, io.ErrShortBuffer
	}

	b := p[:sha1BlockSizeInBytes]

	emptySHA1Block(b)
	l, err := r.r.Read(b)
	r.size += uint64(l * 8)

	if err != nil && err != io.EOF {
		return l, err
	}

	switch {
	case l == 0:
		if r.size%sha1BlockSize == 0 {
			p[l] = 0x80
		}
		r.addMessageLength(p)
		return sha1BlockSizeInBytes, io.EOF
	case l == sha1BlockSizeInBytes:
		return sha1BlockSizeInBytes, nil
	case l < (sha1BlockSizeInBytes - 9):
		p[l] = 0x80
		r.addMessageLength(p)
		return sha1BlockSizeInBytes, io.EOF
	case l < sha1BlockSizeInBytes:
		p[l] = 0x80
		return sha1BlockSizeInBytes, nil
	}

	return sha1BlockSizeInBytes, io.EOF
}
