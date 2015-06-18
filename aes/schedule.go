package aes

func (s keySchedule128) round(i int) roundSchedule {
	return s[i]
}

func scheduleFor(key Key) keySchedule {
	// TODO: implement
	return keySchedule128{}
}
