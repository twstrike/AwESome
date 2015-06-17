package aes

import(
	"bytes"
    "encoding/binary"
	"encoding/hex"
)

func bytesToWord(inp []byte, data interface{}) {
	buf := bytes.NewReader(inp)
	binary.Read(buf, binary.BigEndian, data)
}

func hexStringToWord(in HexString, data interface{}) {
	decodes, _ := hex.DecodeString(string(in))
	bytesToWord(decodes, data)
}
