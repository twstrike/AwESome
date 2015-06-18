package aes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
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

func hexStringToWord(in HexString, data interface{}) {
	decodes, _ := hex.DecodeString(string(in))
	bytesToWord(decodes, data)
}

func wordToHexString(in interface{}) HexString {
	encoded := hex.EncodeToString(wordsToBytes(in))
	return HexString(encoded)
}
