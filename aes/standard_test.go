package aes_test

import (
	"github.com/twstrike/AwESome/aes"
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type GladmanAESSuite struct{}

var _ = Suite(&GladmanAESSuite{})

type StandardAESSuite struct{}

var _ = Suite(&StandardAESSuite{})

type SimpleSuite struct{}

var _ = Suite(&SimpleSuite{})

func (s *SimpleSuite) SetUpSuite(c *C) {
	c.Skip("Not yet implemented.")
}

func (s *StandardAESSuite) SetUpSuite(c *C) {
	c.Skip("Not yet implemented.")
}

func (s *GladmanAESSuite) SetUpSuite(c *C) {
	c.Skip("Not yet implemented.")
}

var testGladman128Encrypting = []string{
	"Plaintext:      32 43 f6 a8 88 5a 30 8d 31 31 98 a2 e0 37 07 34",
	"Key:            2b 7e 15 16 28 ae d2 a6 ab f7 15 88 09 cf 4f 3c",
	"Start round  1: 19 3d e3 be a0 f4 e2 2b 9a c6 8d 2a e9 f8 48 08",
	"Start round  2: a4 9c 7f f2 68 9f 35 2b 6b 5b ea 43 02 6a 50 49",
	"Start round  3: aa 8f 5f 03 61 dd e3 ef 82 d2 4a d2 68 32 46 9a",
	"Start round  4: 48 6c 4e ee 67 1d 9d 0d 4d e3 b1 38 d6 5f 58 e7",
	"Start round  5: e0 92 7f e8 c8 63 63 c0 d9 b1 35 50 85 b8 be 01",
	"Start round  6: f1 00 6f 55 c1 92 4c ef 7c c8 8b 32 5d b5 d5 0c",
	"Start round  7: 26 0e 2e 17 3d 41 b7 7d e8 64 72 a9 fd d2 8b 25",
	"Start round  8: 5a 41 42 b1 19 49 dc 1f a3 e0 19 65 7a 8c 04 0c",
	"Start round  9: ea 83 5c f0 04 45 33 2d 65 5d 98 ad 85 96 b0 c5",
	"Start round 10: eb 40 f2 1e 59 2e 38 84 8b a1 13 e7 1b c3 42 d2",
	"Ciphertext:     39 25 84 1d 02 dc 09 fb dc 11 85 97 19 6a 0b 32",
}

var testGladman128Decrypting = []string{
	"Ciphertext:     39 25 84 1d 02 dc 09 fb dc 11 85 97 19 6a 0b 32",
	"Key:            2b 7e 15 16 28 ae d2 a6 ab f7 15 88 09 cf 4f 3c",
	"Start round  1: e9 31 7d b5 cb 32 2c 72 3d 2e 89 5f af 09 07 94",
	"Start round  2: 87 6e 46 a6 f2 4c e7 8c 4d 90 4a d8 97 ec c3 95",
	"Start round  3: be 3b d4 fe d4 e1 f2 c8 0a 64 2c c0 da 83 86 4d",
	"Start round  4: f7 83 40 3f 27 43 3d f0 9b b5 31 ff 54 ab a9 d3",
	"Start round  5: a1 4f 3d fe 78 e8 03 fc 10 d5 a8 df 4c 63 29 23",
	"Start round  6: e1 fb 96 7c e8 c8 ae 9b 35 6c d2 ba 97 4f fb 53",
	"Start round  7: 52 a4 c8 94 85 11 6a 28 e3 cf 2f d7 f6 50 5e 07",
	"Start round  8: ac c1 d6 b8 ef b5 5a 7b 13 23 cf df 45 73 11 b5",
	"Start round  9: 49 db 87 3b 45 39 53 89 7f 02 d2 f1 77 de 96 1a",
	"Start round 10: d4 bf 5d 30 e0 b4 52 ae b8 41 11 f1 1e 27 98 e5",
	"Plaintext:      32 43 f6 a8 88 5a 30 8d 31 31 98 a2 e0 37 07 34",
}

var testGladman192Encrypting = []string{
	"Plaintext:      32 43 f6 a8 88 5a 30 8d 31 31 98 a2 e0 37 07 34",
	"Key:            2b 7e 15 16 28 ae d2 a6 ab f7 15 88 09 cf 4f 3c 76 2e 71 60 f3 8b 4d a5",
	"Start round  1: 19 3d e3 be a0 f4 e2 2b 9a c6 8d 2a e9 f8 48 08",
	"Start round  2: 72 48 f0 85 13 40 54 3f 5f 65 c0 61 17 35 e7 f1",
	"Start round  3: 14 e2 0a 1f b3 dc 3a 62 36 27 2f d3 da 75 6f 70",
	"Start round  4: cb 42 fd 92 33 3f 28 43 21 11 fe 84 3c bc a8 1a",
	"Start round  5: 94 99 c6 ee b9 78 94 12 bb 04 09 b7 a7 97 c0 25",
	"Start round  6: 8a 6c 1e 3e db 78 a6 4e f5 db 78 62 ea d6 a4 01",
	"Start round  7: 43 5c e2 58 97 7c 16 d8 71 7c 0f f7 79 19 e5 19",
	"Start round  8: 70 b8 37 b9 ae fc 8b bc 5c d2 ab a5 cc 56 d7 4e",
	"Start round  9: 94 a2 c3 31 ed 28 bf de d7 d6 c5 83 4b a9 ed 1e",
	"Start round 10: 52 2d 88 c5 ed ab 19 4e 25 ec 73 1c 11 fa 6b 08",
	"Start round 11: ab 82 54 06 da 72 4d 0c 2b cc f6 c2 39 32 12 01",
	"Start round 12: 43 88 b3 26 6a f7 68 e8 4f cc a4 2a 3a 4d 45 5f",
	"Ciphertext:     f9 fb 29 ae fc 38 4a 25 03 40 d8 33 b8 7e bc 00",
}

var testGladman192Decrypting = []string{
	"Ciphertext:     f9 fb 29 ae fc 38 4a 25 03 40 d8 33 b8 7e bc 00",
	"Key:            2b 7e 15 16 28 ae d2 a6 ab f7 15 88 09 cf 4f 3c 76 2e 71 60 f3 8b 4d a5",
	"Start round  1: 1a 68 49 cf 02 4b 6e f7 84 e3 6d 9b 80 c4 45 e5",
	"Start round  2: 62 40 42 7c 57 4b c9 6f f1 23 20 fe 12 13 e3 25",
	"Start round  3: 00 62 8f 30 55 ce 7f a6 3f 2d c4 2f 82 d8 d4 9c",
	"Start round  4: 22 34 a6 72 55 f6 55 c7 0e d3 2e 1d b3 3a 08 ec",
	"Start round  5: 51 b0 62 2f e4 b5 0e 56 4a b1 9a 65 4b 6c 3d 06",
	"Start round  6: 1a 10 76 d4 88 10 d9 6a a3 d4 98 61 b6 4a 47 68",
	"Start round  7: 7e bc bc 7c b9 b9 49 b2 e6 f6 72 2f 87 50 24 aa",
	"Start round  8: 22 bc 01 3f 56 f2 ba 28 ea 88 b4 c9 5c ee 22 a9",
	"Start round  9: 1f 75 bb a2 c3 82 c2 4f fd 65 54 1a eb 2c 34 5f",
	"Start round 10: fa 86 15 51 6d cc a8 c0 05 9d 67 aa 57 98 80 66",
	"Start round 11: 40 09 ba a1 7d 4d 94 97 cf 96 8c 75 f0 52 20 ef",
	"Start round 12: d4 bf 5d 30 e0 b4 52 ae b8 41 11 f1 1e 27 98 e5",
	"Plaintext:      32 43 f6 a8 88 5a 30 8d 31 31 98 a2 e0 37 07 34",
}

var testGladman256Encrypting = []string{
	"Plaintext:      32 43 f6 a8 88 5a 30 8d 31 31 98 a2 e0 37 07 34",
	"Key:            2b 7e 15 16 28 ae d2 a6 ab f7 15 88 09 cf 4f 3c 76 2e 71 60 f3 8b 4d a5 6a 78 4d 90 45 19 0c fe",
	"Start round  1: 19 3d e3 be a0 f4 e2 2b 9a c6 8d 2a e9 f8 48 08",
	"Start round  2: 72 48 f0 85 13 40 54 3f 22 80 9e ea 6d 1f 2a b2",
	"Start round  3: 59 f8 a8 d4 12 71 bf 44 e2 2b a6 5e 5d 69 9a 49",
	"Start round  4: 88 a8 eb d5 66 49 40 5e 9f ad 55 e9 33 0d 7f 84",
	"Start round  5: 6d 0c 80 51 d5 bc 1d b5 c5 1f 45 0f 18 46 7f 34",
	"Start round  6: cf f5 04 43 27 6f 76 55 a5 5a fd 7b b6 99 f4 5f",
	"Start round  7: 63 93 d6 68 2d da 2e 4f 42 88 77 37 12 57 8a 11",
	"Start round  8: 29 26 ae 58 f4 32 23 4b f0 70 ff 6e 56 9e 44 23",
	"Start round  9: 3f da e4 32 0e ce 55 ce c9 32 d6 55 e4 3a cd 2b",
	"Start round 10: d6 40 90 18 38 64 11 35 61 ef 7c 37 99 00 31 fd",
	"Start round 11: 6b 3c e6 72 ea e1 1d 52 e2 8f 1d 54 96 e0 c0 d0",
	"Start round 12: 84 74 88 72 49 e5 0a 9f 17 c0 5a 37 a1 a6 9f 41",
	"Start round 13: bf 8a 29 14 80 f8 06 21 44 3e 2b 81 aa 2f 4c 16",
	"Start round 14: d3 20 3d d1 1a c7 2d 8b 5e c4 72 24 95 3d fe 5b",
	"Ciphertext:     1a 6e 6c 2c 66 2e 7d a6 50 1f fb 62 bc 9e 93 f3",
}

var testGladman256Decrypting = []string{
	"Ciphertext:     1a 6e 6c 2c 66 2e 7d a6 50 1f fb 62 bc 9e 93 f3",
	"Key:            2b 7e 15 16 28 ae d2 a6 ab f7 15 88 09 cf 4f 3c 76 2e 71 60 f3 8b 4d a5 6a 78 4d 90 45 19 0c fe",
	"Start round  1: 66 c6 40 39 a2 1c bb 3e 58 27 27 3d 2a b7 d8 36",
	"Start round  2: 08 41 f1 47 cd b2 29 fa 1b 15 a5 fd ac 7e 6f 0c",
	"Start round  3: 5f d9 be 83 3b ba db 40 f0 24 c4 db 32 92 67 9a",
	"Start round  4: 7f f8 a4 70 87 73 ba 40 98 e1 8e 00 90 eb a4 20",
	"Start round  5: f6 43 10 54 07 df c7 ad ef 63 60 96 ee 09 82 9a",
	"Start round  6: 75 8b f6 f1 ab 23 bd 23 dd 80 69 8b 69 57 fc fc",
	"Start round  7: a5 23 16 26 bf 51 1b 6a 8c 0b e4 b3 b1 f7 26 9f",
	"Start round  8: fb 57 f5 82 d8 c4 7e 45 2c 5b f6 84 c9 dc 31 9a",
	"Start round  9: 8a a8 54 cf cc be bf 1a 06 ee f2 fc 4e e6 38 21",
	"Start round 10: 3c 65 6e 18 03 c0 d2 d1 a6 5a cd d5 ad fe a4 76",
	"Start round 11: c4 3b fc 5f 33 95 d2 03 db d7 e9 58 c3 c2 09 1e",
	"Start round 12: cb a3 24 3b c9 f1 b8 48 98 f9 c2 1b 4c 41 08 58",
	"Start round 13: 40 09 0b 37 7d cd e5 97 93 c0 8c 75 3c 52 20 87",
	"Start round 14: d4 bf 5d 30 e0 b4 52 ae b8 41 11 f1 1e 27 98 e5",
	"Plaintext:      32 43 f6 a8 88 5a 30 8d 31 31 98 a2 e0 37 07 34",
}

var testAES128Encrypting = []string{
	"Plaintext:      00 11 22 33 44 55 66 77 88 99 aa bb cc dd ee ff",
	"Key:            00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f",
	"Start round  1: 00 10 20 30 40 50 60 70 80 90 a0 b0 c0 d0 e0 f0",
	"Start round  2: 89 d8 10 e8 85 5a ce 68 2d 18 43 d8 cb 12 8f e4",
	"Start round  3: 49 15 59 8f 55 e5 d7 a0 da ca 94 fa 1f 0a 63 f7",
	"Start round  4: fa 63 6a 28 25 b3 39 c9 40 66 8a 31 57 24 4d 17",
	"Start round  5: 24 72 40 23 69 66 b3 fa 6e d2 75 32 88 42 5b 6c",
	"Start round  6: c8 16 77 bc 9b 7a c9 3b 25 02 79 92 b0 26 19 96",
	"Start round  7: c6 2f e1 09 f7 5e ed c3 cc 79 39 5d 84 f9 cf 5d",
	"Start round  8: d1 87 6c 0f 79 c4 30 0a b4 55 94 ad d6 6f f4 1f",
	"Start round  9: fd e3 ba d2 05 e5 d0 d7 35 47 96 4e f1 fe 37 f1",
	"Start round 10: bd 6e 7c 3d f2 b5 77 9e 0b 61 21 6e 8b 10 b6 89",
	"Ciphertext:     69 c4 e0 d8 6a 7b 04 30 d8 cd b7 80 70 b4 c5 5a",
}

var testAES128Decrypting = []string{
	"Ciphertext:     69 c4 e0 d8 6a 7b 04 30 d8 cd b7 80 70 b4 c5 5a",
	"Key:            00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f",
	"Start round  1: 7a d5 fd a7 89 ef 4e 27 2b ca 10 0b 3d 9f f5 9f",
	"Start round  2: 54 d9 90 a1 6b a0 9a b5 96 bb f4 0e a1 11 70 2f",
	"Start round  3: 3e 1c 22 c0 b6 fc bf 76 8d a8 50 67 f6 17 04 95",
	"Start round  4: b4 58 12 4c 68 b6 8a 01 4b 99 f8 2e 5f 15 55 4c",
	"Start round  5: e8 da b6 90 14 77 d4 65 3f f7 f5 e2 e7 47 dd 4f",
	"Start round  6: 36 33 9d 50 f9 b5 39 26 9f 2c 09 2d c4 40 6d 23",
	"Start round  7: 2d 6d 7e f0 3f 33 e3 34 09 36 02 dd 5b fb 12 c7",
	"Start round  8: 3b d9 22 68 fc 74 fb 73 57 67 cb e0 c0 59 0e 2d",
	"Start round  9: a7 be 1a 69 97 ad 73 9b d8 c9 ca 45 1f 61 8b 61",
	"Start round 10: 63 53 e0 8c 09 60 e1 04 cd 70 b7 51 ba ca d0 e7",
	"Plaintext:      00 11 22 33 44 55 66 77 88 99 aa bb cc dd ee ff",
}

var testAES192Encrypting = []string{
	"Plaintext:      00 11 22 33 44 55 66 77 88 99 aa bb cc dd ee ff",
	"Key:            00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17",
	"Start round  1: 00 10 20 30 40 50 60 70 80 90 a0 b0 c0 d0 e0 f0",
	"Start round  2: 4f 63 76 06 43 e0 aa 85 af f8 c9 d0 41 fa 0d e4",
	"Start round  3: cb 02 81 8c 17 d2 af 9c 62 aa 64 42 8b b2 5f d7",
	"Start round  4: f7 5c 77 78 a3 27 c8 ed 8c fe bf c1 a6 c3 7f 53",
	"Start round  5: 22 ff c9 16 a8 14 74 41 64 96 f1 9c 64 ae 25 32",
	"Start round  6: 80 12 1e 07 76 fd 1d 8a 8d 8c 31 bc 96 5d 1f ee",
	"Start round  7: 67 1e f1 fd 4e 2a 1e 03 df dc b1 ef 3d 78 9b 30",
	"Start round  8: 0c 03 70 d0 0c 01 e6 22 16 6b 8a cc d6 db 3a 2c",
	"Start round  9: 72 55 da d3 0f b8 03 10 e0 0d 6c 6b 40 d0 52 7c",
	"Start round 10: a9 06 b2 54 96 8a f4 e9 b4 bd b2 d2 f0 c4 43 36",
	"Start round 11: 88 ec 93 0e f5 e7 e4 b6 cc 32 f4 c9 06 d2 94 14",
	"Start round 12: af b7 3e eb 1c d1 b8 51 62 28 0f 27 fb 20 d5 85",
	"Ciphertext:     dd a9 7c a4 86 4c df e0 6e af 70 a0 ec 0d 71 91",
}

var testAES192Decrypting = []string{
	"Ciphertext:     dd a9 7c a4 86 4c df e0 6e af 70 a0 ec 0d 71 91",
	"Key:            00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17",
	"Start round  1: 79 3e 76 97 9c 34 03 e9 aa b7 b2 d1 0f a9 6c cc",
	"Start round  2: c4 94 bf fa e6 23 22 ab 4b b5 dc 4e 6f ce 69 dd",
	"Start round  3: d3 7e 37 05 90 7a 1a 20 8d 1c 37 1e 8c 6f bf b5",
	"Start round  4: 40 6c 50 10 76 d7 00 66 e1 70 57 ca 09 fc 7b 7f",
	"Start round  5: fe 7c 7e 71 fe 7f 80 70 47 b9 51 93 f6 7b 8e 4b",
	"Start round  6: 85 e5 c8 04 2f 86 14 54 9e bc a1 7b 27 72 72 df",
	"Start round  7: cd 54 c7 28 38 64 c0 c5 5d 4c 72 7e 90 c9 a4 65",
	"Start round  8: 93 fa a1 23 c2 90 3f 47 43 e4 dd 83 43 16 92 de",
	"Start round  9: 68 cc 08 ed 0a bb d2 bc 64 2e f5 55 24 4a e8 78",
	"Start round 10: 1f b5 43 0e f0 ac cf 64 aa 37 0c de 3d 77 79 2c",
	"Start round 11: 84 e1 dd 69 1a 41 d7 6f 79 2d 38 97 83 fb ac 70",
	"Start round 12: 63 53 e0 8c 09 60 e1 04 cd 70 b7 51 ba ca d0 e7",
	"Plaintext:      00 11 22 33 44 55 66 77 88 99 aa bb cc dd ee ff",
}

var testAES256Encrypting = []string{
	"Plaintext:      00 11 22 33 44 55 66 77 88 99 aa bb cc dd ee ff",
	"Key:            00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f",
	"Start round  1: 00 10 20 30 40 50 60 70 80 90 a0 b0 c0 d0 e0 f0",
	"Start round  2: 4f 63 76 06 43 e0 aa 85 ef a7 21 32 01 a4 e7 05",
	"Start round  3: 18 59 fb c2 8a 1c 00 a0 78 ed 8a ad c4 2f 61 09",
	"Start round  4: 97 5c 66 c1 cb 9f 3f a8 a9 3a 28 df 8e e1 0f 63",
	"Start round  5: 1c 05 f2 71 a4 17 e0 4f f9 21 c5 c1 04 70 15 54",
	"Start round  6: c3 57 aa e1 1b 45 b7 b0 a2 c7 bd 28 a8 dc 99 fa",
	"Start round  7: 7f 07 41 43 cb 4e 24 3e c1 0c 81 5d 83 75 d5 4c",
	"Start round  8: d6 53 a4 69 6c a0 bc 0f 5a ca ab 5d b9 6c 5e 7d",
	"Start round  9: 5a a8 58 39 5f d2 8d 7d 05 e1 a3 88 68 f3 b9 c5",
	"Start round 10: 4a 82 48 51 c5 7e 7e 47 64 3d e5 0c 2a f3 e8 c9",
	"Start round 11: c1 49 07 f6 ca 3b 3a a0 70 e9 aa 31 3b 52 b5 ec",
	"Start round 12: 5f 9c 6a bf ba c6 34 aa 50 40 9f a7 66 67 76 53",
	"Start round 13: 51 66 04 95 43 53 95 03 14 fb 86 e4 01 92 25 21",
	"Start round 14: 62 7b ce b9 99 9d 5a aa c9 45 ec f4 23 f5 6d a5",
	"Ciphertext:     8e a2 b7 ca 51 67 45 bf ea fc 49 90 4b 49 60 89",
}

var testAES256Decrypting = []string{
	"Ciphertext:     8e a2 b7 ca 51 67 45 bf ea fc 49 90 4b 49 60 89",
	"Key:            00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f",
	"Start round  1: aa 5e ce 06 ee 6e 3c 56 dd e6 8b ac 26 21 be bf",
	"Start round  2: d1 ed 44 fd 1a 0f 3f 2a fa 4f f2 7b 7c 33 2a 69",
	"Start round  3: cf b4 db ed f4 09 38 08 53 85 02 ac 33 de 18 5c",
	"Start round  4: 78 e2 ac ce 74 1e d5 42 51 00 c5 e0 e2 3b 80 c7",
	"Start round  5: d6 f3 d9 dd a6 27 9b d1 43 0d 52 a0 e5 13 f3 fe",
	"Start round  6: be b5 0a a6 cf f8 56 12 6b 0d 6a ff 45 c2 5d c4",
	"Start round  7: f6 e0 62 ff 50 74 58 f9 be 50 49 76 56 ed 65 4c",
	"Start round  8: d2 2f 0c 29 1f fe 03 1a 78 9d 83 b2 ec c5 36 4c",
	"Start round  9: 2e 6e 7a 2d af c6 ee f8 3a 86 ac e7 c2 5b a9 34",
	"Start round 10: 9c f0 a6 20 49 fd 59 a3 99 51 89 84 f2 6b e1 78",
	"Start round 11: 88 db 34 fb 1f 80 76 78 d3 f8 33 c2 19 4a 75 9e",
	"Start round 12: ad 9c 7e 01 7e 55 ef 25 bc 15 0f e0 1c cb 63 95",
	"Start round 13: 84 e1 fd 6b 1a 5c 94 6f df 49 38 97 7c fb ac 23",
	"Start round 14: 63 53 e0 8c 09 60 e1 04 cd 70 b7 51 ba ca d0 e7",
	"Plaintext:      00 11 22 33 44 55 66 77 88 99 aa bb cc dd ee ff",
}

var testAllZeroes128Encrypting = []string{
	"Plaintext:      00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
	"Key:            00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
	"Start round  1: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
	"Start round  2: 01 00 00 00 01 00 00 00 01 00 00 00 01 00 00 00",
	"Start round  3: c6 e4 e4 8b a4 87 87 e8 c6 e4 e4 8b a4 87 87 e8",
	"Start round  4: 28 2d f3 c4 6a f3 86 25 4a 4e 90 a7 08 90 e5 46",
	"Start round  5: ab d2 cd fe 37 5a b5 49 50 a0 af c0 75 9a 6a 5f",
	"Start round  6: d4 6f 4f 6c 55 b8 96 33 7e 05 bb 3d 79 79 de 23",
	"Start round  7: 04 f2 ca 97 07 78 28 45 e2 2f 01 96 49 c5 d7 10",
	"Start round  8: b7 aa e4 c5 1d 25 2d 4f 6c 92 0f 81 94 e5 81 50",
	"Start round  9: 23 e7 8c 3c 13 21 63 db aa c0 c6 57 2e 03 cb 95",
	"Start round 10: 7f fe 0e 95 51 a5 66 35 0e 34 7c 47 29 29 ec cb",
	"Ciphertext:     66 e9 4b d4 ef 8a 2c 3b 88 4c fa 59 ca 34 2b 2e",
}

var testAllZeroes128Decrypting = []string{
	"Ciphertext:     66 e9 4b d4 ef 8a 2c 3b 88 4c fa 59 ca 34 2b 2e",
	"Key:            00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
	"Start round  1: d2 06 10 1f d1 18 ce 2a ab a5 ab 96 a5 bb 33 a0",
	"Start round  2: 26 fd b4 2a 7d ba 1f eb ac 7b 64 b9 31 94 fb 5b",
	"Start round  3: a9 3f 76 53 a4 4f 0c a6 50 d9 69 84 22 ac d8 0c",
	"Start round  4: f2 bc 7c ca c5 15 0e 88 98 a6 74 6e 3b 89 34 90",
	"Start round  5: 48 6c ea 26 fc 6b 1d 50 f3 b6 84 c3 b6 a8 90 27",
	"Start round  6: 62 be 79 cf 9a e0 02 bb 53 b8 bd 3b 9d b5 d5 ba",
	"Start round  7: 34 0d 60 5a 02 2f d9 1c d6 60 0d 3f 30 d8 44 5c",
	"Start round  8: b4 17 69 9b 49 69 17 3d b4 17 69 9b 49 69 17 3d",
	"Start round  9: 7c 63 63 63 7c 63 63 63 7c 63 63 63 7c 63 63 63",
	"Start round 10: 63 63 63 63 63 63 63 63 63 63 63 63 63 63 63 63",
	"Plaintext:      00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
}

var testKeyOne128Encrypting = []string{
	"Plaintext:      00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
	"Key:            00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01",
	"Start round  1: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01",
	"Start round  2: 1e 1f 3e 3e 01 00 1f 00 01 00 1f 00 01 00 1f 01",
	"Start round  3: 66 e0 d8 04 d6 43 e3 f2 17 67 13 5a f9 cf 28 e9",
	"Start round  4: 7e 41 45 5e 09 86 08 b5 a1 69 f0 da 70 61 e1 bf",
	"Start round  5: 79 89 5e 0d d2 19 17 93 a2 96 41 74 c2 e1 0d 64",
	"Start round  6: a0 f7 36 25 ca ae 92 22 a4 76 49 4b 04 d0 d6 8b",
	"Start round  7: a8 1b dd b6 3b f7 72 10 81 6b 51 a9 70 27 04 e1",
	"Start round  8: 89 51 a6 38 37 41 9e 27 9c 5d fe a7 d0 20 3c 26",
	"Start round  9: 4c 39 fa 23 09 9f 8d 0b 53 f7 13 4f 0a 53 02 53",
	"Start round 10: dd 9d a7 03 9c a1 f1 58 42 43 46 94 5d c8 68 7a",
	"Ciphertext:     05 45 aa d5 6d a2 a9 7c 36 63 d1 43 2a 3d 1c 84",
}

var testKeyOne128Decrypting = []string{
	"Ciphertext:     05 45 aa d5 6d a2 a9 7c 36 63 d1 43 2a 3d 1c 84",
	"Key:            00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01",
	"Start round  1: c1 32 5a da de 1a 45 7b 2c e8 5c 6a 4c 5e a1 22",
	"Start round  2: 29 db 7d ed 01 68 77 26 ed ed 2d 2b 67 12 5d 84",
	"Start round  3: a7 83 bb f7 9a 4c eb 07 de b7 24 cc 70 d1 0b 5c",
	"Start round  4: c2 68 d1 f8 e2 7f f2 4e 0c cc c1 ca 51 af 40 d3",
	"Start round  5: e0 e4 3b 3d 74 38 f6 3f 49 70 05 93 f2 68 4f b3",
	"Start round  6: b6 d4 83 43 b5 90 d7 d7 3a f8 58 dc 25 a7 f0 92",
	"Start round  7: f3 44 8c 08 01 f9 f8 58 32 ef 6e d5 51 83 30 57",
	"Start round  8: 33 1a 7d 1e f6 85 34 f2 f0 8a 61 89 99 e1 11 be",
	"Start round  9: 72 63 c0 7c 7c 63 c0 b2 7c 63 b2 63 7c c0 c0 63",
	"Start round 10: 63 63 63 7c 63 63 63 63 63 63 63 63 63 63 63 63",
	"Plaintext:      00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
}

var testPlainOne128Encrypting = []string{
	"Plaintext:      00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01",
	"Key:            00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
	"Start round  1: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01",
	"Start round  2: 1e 1f 21 3e 01 00 00 00 01 00 00 00 01 00 00 00",
	"Start round  3: da ea ea 99 75 56 ef 51 58 5d c3 15 5a da 24 4b",
	"Start round  4: 6b 78 40 4b 44 dd 5b fb 89 73 17 e4 78 87 db aa",
	"Start round  5: 14 47 70 88 31 17 26 e7 1f 03 f1 40 92 6c 27 a3",
	"Start round  6: 30 dd 66 d8 e8 fe 56 7f 23 dd e9 d2 68 54 93 04",
	"Start round  7: de d8 c5 d3 dc 39 2f e5 54 1e 04 f1 bc 15 a2 69",
	"Start round  8: 26 b8 f7 a9 e8 1a 48 1e 78 15 35 0a 1b 22 26 19",
	"Start round  9: 44 24 06 b6 d9 51 38 9d 40 4d 63 c4 95 fb 34 85",
	"Start round 10: 83 f7 5d d0 88 f3 8d e2 2f 83 89 19 59 35 7f 36",
	"Ciphertext:     58 e2 fc ce fa 7e 30 61 36 7f 1d 57 a4 e7 45 5a",
}

var testPlainOne128Decrypting = []string{
	"Ciphertext:     58 e2 fc ce fa 7e 30 61 36 7f 1d 57 a4 e7 45 5a",
	"Key:            00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00",
	"Start round  1: ec 0d a7 05 c4 ec d2 70 15 96 4c 98 cb 68 5d d4",
	"Start round  2: 1b d1 fb 97 35 e3 18 4e 09 0f 6f 5e 2a 36 07 1c",
	"Start round  3: f7 a2 96 d4 9b 59 f7 d3 bc 93 68 72 af 6c 52 67",
	"Start round  4: 1d 12 f2 f9 86 72 3a 66 20 59 a6 d9 65 61 15 a1",
	"Start round  5: 04 bb 1e f2 9b c1 dc 61 26 20 33 d2 45 c1 b1 b5",
	"Start round  6: fa f0 a1 0a c7 7b cc c4 c0 50 51 94 4f a0 f7 09",
	"Start round  7: 7f c1 f0 ac 1b 8f b9 b3 a7 17 09 0f bc bc 39 69",
	"Start round  8: 57 b1 2e b3 9d 4c 36 ee 6a 57 87 d1 be 87 df 59",
	"Start round  9: 72 63 63 63 7c 63 63 b2 7c 63 fd 63 7c c0 63 63",
	"Start round 10: 63 63 63 7c 63 63 63 63 63 63 63 63 63 63 63 63",
	"Plaintext:      00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01",
}

type EncryptionTest struct {
	plaintext, key, expected string
}

type DecryptionTest struct {
	ciphertext, key, expected string
}

func extractDataFrom(data string) string {
	return strings.Replace(data[16:len(data)], " ", "", -1)
}

func encrypt(plain, key string) string {
	return string(aes.EncryptHex(aes.HexString(key), aes.HexString(plain)))
}

func decrypt(cipher, key string) string {
	return ""
}

func parseEncryptionData(data []string) EncryptionTest {
	plain := extractDataFrom(data[0])
	key := extractDataFrom(data[1])
	expected := extractDataFrom(data[len(data)-1])
	return EncryptionTest{plain, key, expected}
}

func parseDecryptionData(data []string) DecryptionTest {
	cipher := extractDataFrom(data[0])
	key := extractDataFrom(data[1])
	expected := extractDataFrom(data[len(data)-1])
	return DecryptionTest{cipher, key, expected}
}

func testEncryptingOn(data []string, c *C) {
	encryptionData := parseEncryptionData(data)
	result := encrypt(encryptionData.plaintext, encryptionData.key)
	c.Assert(result, Equals, encryptionData.expected)
}

func testDecryptingOn(data []string, c *C) {
	decryptionData := parseDecryptionData(data)
	result := decrypt(decryptionData.ciphertext, decryptionData.key)
	c.Assert(result, Equals, decryptionData.expected)
}

func (s *GladmanAESSuite) TestGladman128Encrypting(c *C) {
	testEncryptingOn(testGladman128Encrypting, c)
}

func (s *GladmanAESSuite) TestGladman128Decrypting(c *C) {
	testDecryptingOn(testGladman128Decrypting, c)
}

func (s *GladmanAESSuite) TestGladman192Encrypting(c *C) {
	testEncryptingOn(testGladman192Encrypting, c)
}

func (s *GladmanAESSuite) TestGladman192Decrypting(c *C) {
	testDecryptingOn(testGladman192Decrypting, c)
}

func (s *GladmanAESSuite) TestGladman256Encrypting(c *C) {
	testEncryptingOn(testGladman256Encrypting, c)
}

func (s *GladmanAESSuite) TestGladman256Decrypting(c *C) {
	testDecryptingOn(testGladman256Decrypting, c)
}

func (s *StandardAESSuite) TestStandard128Encrypting(c *C) {
	testEncryptingOn(testAES128Encrypting, c)
}

func (s *StandardAESSuite) TestStandard128Decrypting(c *C) {
	testDecryptingOn(testAES128Decrypting, c)
}

func (s *StandardAESSuite) TestStandard192Encrypting(c *C) {
	testEncryptingOn(testAES192Encrypting, c)
}

func (s *StandardAESSuite) TestStandard192Decrypting(c *C) {
	testDecryptingOn(testAES192Decrypting, c)
}

func (s *StandardAESSuite) TestStandard256Encrypting(c *C) {
	testEncryptingOn(testAES256Encrypting, c)
}

func (s *StandardAESSuite) TestStandard256Decrypting(c *C) {
	testDecryptingOn(testAES256Decrypting, c)
}

func (s *SimpleSuite) TestAllZeroes128Encrypting(c *C) {
	testEncryptingOn(testAllZeroes128Encrypting, c)
}

func (s *SimpleSuite) TestAllZeroes128Decrypting(c *C) {
	testDecryptingOn(testAllZeroes128Decrypting, c)
}

func (s *SimpleSuite) TestKeyOne128Encrypting(c *C) {
	testEncryptingOn(testKeyOne128Encrypting, c)
}

func (s *SimpleSuite) TestKeyOne128Decrypting(c *C) {
	testDecryptingOn(testKeyOne128Decrypting, c)
}

func (s *SimpleSuite) TestPlainOne128Encrypting(c *C) {
	testEncryptingOn(testPlainOne128Encrypting, c)
}

func (s *SimpleSuite) TestPlainOne128Decrypting(c *C) {
	testDecryptingOn(testPlainOne128Decrypting, c)
}
