package aes

const Nr128 = 10
const Nr192 = 12
const Nr256 = 14

type roundSchedule [4]word
type keySchedule128 [Nr128 + 1]roundSchedule

func scheduleFor(key Key128) keySchedule128 {
	return keySchedule128{}
}
