package hash

import (
	"bytes"
	"encoding/hex"
	"io"

	. "gopkg.in/check.v1"
)

type SHA1MessageReaderSuite struct{}

var _ = Suite(&SHA1MessageReaderSuite{})

func (s *SHA1MessageReaderSuite) TestPaddingNecessary(c *C) {
	input, _ := hex.DecodeString("0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"7172737475767778")
	reader := NewSha1MessageReader(bytes.NewBuffer(input))

	buffer := sha1Block{}
	n, err := reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48,
		0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58,
		0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68,
		0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78,
	})

	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, nil)

	buffer = sha1Block{}
	n, err = reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00,
	})

	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, io.EOF)
}

func (s *SHA1MessageReaderSuite) TestPaddingSingleBlock(c *C) {
	input, _ := hex.DecodeString("6162636465")
	reader := NewSha1MessageReader(bytes.NewBuffer(input))

	buffer := sha1Block{}
	n, err := reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x61, 0x62, 0x63, 0x64, 0x65, 0x80, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x28,
	})

	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, io.EOF)
}

func (s *SHA1MessageReaderSuite) TestPaddingSecondBlock(c *C) {
	input, _ := hex.DecodeString("0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"7172737475767778" +
		"0102")
	reader := NewSha1MessageReader(bytes.NewBuffer(input))

	buffer := sha1Block{}
	n, err := reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48,
		0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58,
		0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68,
		0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78,
	})
	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, nil)

	buffer = sha1Block{}
	n, err = reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x01, 0x02, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x10,
	})
	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, io.EOF)
}

func (s *SHA1MessageReaderSuite) TestPaddingBlockSizeLimit(c *C) {
	input, _ := hex.DecodeString("0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"71727374757677")
	reader := NewSha1MessageReader(bytes.NewBuffer(input))

	buffer := sha1Block{}
	n, err := reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48,
		0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58,
		0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68,
		0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x80,
	})
	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, nil)

	buffer = sha1Block{}
	n, err = reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xF8,
	})
	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, io.EOF)
}

func (s *SHA1MessageReaderSuite) TestPaddingInputHavingMultipleBlocks(c *C) {
	input, _ := hex.DecodeString("0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"7172737475767778" +

		"0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"7172737475767778" +

		"0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"71727374757677")

	reader := NewSha1MessageReader(bytes.NewBuffer(input))

	for i := 0; i < 2; i++ {
		buffer := sha1Block{}
		n, err := reader.Read(buffer[:])
		c.Assert(buffer[:], DeepEquals, []byte{
			0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
			0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
			0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
			0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
			0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48,
			0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58,
			0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68,
			0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78,
		})
		c.Assert(n, Equals, sha1BlockSizeInBytes)
		c.Assert(err, Equals, nil)
	}

	var buffer sha1Block
	n, err := reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38,
		0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48,
		0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58,
		0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68,
		0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x80,
	})
	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, nil)

	buffer = sha1Block{}
	n, err = reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0xF8,
	})
	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, io.EOF)
}

func (s *SHA1MessageReaderSuite) TestPaddingNISTcase(c *C) {
	input := []byte("abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq")
	reader := NewSha1MessageReader(bytes.NewBuffer(input))

	buffer := sha1Block{}
	n, err := reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x61, 0x62, 0x63, 0x64,
		0x62, 0x63, 0x64, 0x65,
		0x63, 0x64, 0x65, 0x66,
		0x64, 0x65, 0x66, 0x67,
		0x65, 0x66, 0x67, 0x68,
		0x66, 0x67, 0x68, 0x69,
		0x67, 0x68, 0x69, 0x6a,
		0x68, 0x69, 0x6a, 0x6b,
		0x69, 0x6a, 0x6b, 0x6c,
		0x6a, 0x6b, 0x6c, 0x6d,
		0x6b, 0x6c, 0x6d, 0x6e,
		0x6c, 0x6d, 0x6e, 0x6f,
		0x6d, 0x6e, 0x6f, 0x70,
		0x6e, 0x6f, 0x70, 0x71,
		0x80, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	})

	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, nil)

	n, err = reader.Read(buffer[:])
	c.Assert(buffer[:], DeepEquals, []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0xC0,
	})

	c.Assert(n, Equals, sha1BlockSizeInBytes)
	c.Assert(err, Equals, io.EOF)
}