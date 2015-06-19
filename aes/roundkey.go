package aes

func addRoundKey(s state, rs roundSchedule) state {
	out := state{}
	for i, si := range s {
		column := i % 4
		row := i / 4
		shift := uint(24 - 8*(row))
		word := byte(rs[column] >> shift)
		out[i] = si ^ word
	}
	return out
}
