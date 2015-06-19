package aes

import (
	. "gopkg.in/check.v1"
)

type StandardScheduleSuite struct{}

var _ = Suite(&StandardScheduleSuite{})

func (s *StandardScheduleSuite) Test_scheduleFor128(c *C) {
	key := Key128{0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x09cf4f3c}
	expected := keySchedule128{
		roundSchedule{0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x09cf4f3c},
		roundSchedule{0xa0fafe17, 0x88542cb1, 0x23a33939, 0x2a6c7605},
		roundSchedule{0xf2c295f2, 0x7a96b943, 0x5935807a, 0x7359f67f},
		roundSchedule{0x3d80477d, 0x4716fe3e, 0x1e237e44, 0x6d7a883b},
		roundSchedule{0xef44a541, 0xa8525b7f, 0xb671253b, 0xdb0bad00},
		roundSchedule{0xd4d1c6f8, 0x7c839d87, 0xcaf2b8bc, 0x11f915bc},
		roundSchedule{0x6d88a37a, 0x110b3efd, 0xdbf98641, 0xca0093fd},
		roundSchedule{0x4e54f70e, 0x5f5fc9f3, 0x84a64fb2, 0x4ea6dc4f},
		roundSchedule{0xead27321, 0xb58dbad2, 0x312bf560, 0x7f8d292f},
		roundSchedule{0xac7766f3, 0x19fadc21, 0x28d12941, 0x575c006e},
		roundSchedule{0xd014f9a8, 0xc9ee2589, 0xe13f0cc8, 0xb6630ca6},
	}

	c.Check(scheduleFor(key), DeepEquals, expected)
}

func (s *StandardScheduleSuite) TestRotWord(c *C) {
	c.Check(rotWord(0xABCDEF12), Equals, word(0xCDEF12AB))
}

func (s *StandardScheduleSuite) TestSubWord(c *C) {
	c.Check(subWord(0xABCDEF12), Equals, word(0x62BDDFC9))
}

func (s *StandardScheduleSuite) Test_rcon(c *C) {
	c.Check(rcon(1), Equals, byte(0x01))
	c.Check(rcon(2), Equals, byte(0x02))
	c.Check(rcon(3), Equals, byte(0x04))
	c.Check(rcon(4), Equals, byte(0x08))
	c.Check(rcon(9), Equals, byte(0x1B))
	c.Check(rcon(254), Equals, byte(0xCB))
	c.Check(rcon(255), Equals, byte(0x8D))
}
