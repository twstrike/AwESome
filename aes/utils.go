package aes

import (
	"bytes"
	"encoding/binary"
)

func bytesToWord(inp []byte, data interface{}) {
	buf := bytes.NewReader(inp)
	binary.Read(buf, binary.BigEndian, data)
}

func wordsToBytes(in interface{}) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, in)
	return buffer.Bytes()
}
