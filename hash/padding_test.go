package hash

import (
	"bytes"
	"encoding/hex"

	. "gopkg.in/check.v1"
)

type SHA1PaddingSuite struct{}

var _ = Suite(&SHA1PaddingSuite{})

func (s *SHA1PaddingSuite) TestNoPaddingNecessary(c *C) {
	r, _ := hex.DecodeString("0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"7172737475767778")

	reader := sha1Reader{bytes.NewBuffer(r), 0, false}
	var buffer sha1Block

	hasContent, ended := reader.readWithPadding(&buffer)
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

	c.Assert(hasContent, Equals, true)
	c.Assert(ended, Equals, false)

	hasContent, ended = reader.readWithPadding(&buffer)
	c.Assert(hasContent, Equals, false)
	c.Assert(ended, Equals, true)
}

func (s *SHA1PaddingSuite) TestPaddingSingleBlock(c *C) {
	r, _ := hex.DecodeString("6162636465")
	reader := sha1Reader{bytes.NewBuffer(r), 0, false}
	var buffer sha1Block
	_, result := reader.readWithPadding(&buffer)
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
	c.Assert(result, Equals, true)
}

func (s *SHA1PaddingSuite) TestPaddingSecondBlock(c *C) {
	r, _ := hex.DecodeString("0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"7172737475767778" +
		"0102")
	reader := sha1Reader{bytes.NewBuffer(r), 0, false}
	var buffer sha1Block
	_, result := reader.readWithPadding(&buffer)
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
	c.Assert(result, Equals, false)

	buffer = sha1Block{}
	_, result = reader.readWithPadding(&buffer)
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
	c.Assert(result, Equals, true)
}

func (s *SHA1PaddingSuite) TestPaddingBlockSizeLimit(c *C) {
	r, _ := hex.DecodeString("0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"71727374757677")
	reader := sha1Reader{bytes.NewBuffer(r), 0, false}
	var buffer sha1Block
	_, result := reader.readWithPadding(&buffer)
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
	c.Assert(result, Equals, false)

	buffer = sha1Block{}
	_, result = reader.readWithPadding(&buffer)
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
	c.Assert(result, Equals, true)
}

func (s *SHA1PaddingSuite) TestPaddingInputHavingMultipleBlocks(c *C) {
	b := "0102030405060708" +
		"1112131415161718" +
		"2122232425262728" +
		"3132333435363738" +
		"4142434445464748" +
		"5152535455565758" +
		"6162636465666768" +
		"7172737475767778"

	t := ""
	for i := 0; i < 5; i++ {
		t += b
	}

	r, _ := hex.DecodeString(t[:len(t)-2])
	reader := sha1Reader{bytes.NewBuffer(r), 0, false}

	for i := 0; i < 4; i++ {
		buffer := sha1Block{}
		_, ended := reader.readWithPadding(&buffer)

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
		c.Assert(ended, Equals, false)
	}

	var buffer sha1Block
	_, ended := reader.readWithPadding(&buffer)
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
	c.Assert(ended, Equals, false)

	buffer = sha1Block{}
	_, ended = reader.readWithPadding(&buffer)
	c.Assert(buffer[:], DeepEquals, []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0xF8,
	})
	c.Assert(ended, Equals, true)
}
