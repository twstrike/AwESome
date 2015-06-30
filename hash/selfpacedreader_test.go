package hash

import "io"

type selfPacedReader struct {
	schedule [][]byte
	round    int
}

func newSelfPacedReader(s [][]byte) io.Reader {
	return &selfPacedReader{schedule: s}
}

//FIXME: It doesnt take into consideration when one round cant be copied to p at once
func (r *selfPacedReader) Read(p []byte) (int, error) {
	if r.round > len(r.schedule) {
		return 0, io.EOF
	}

	n := copy(p[:], r.schedule[r.round])
	r.round++
	return n, nil
}
