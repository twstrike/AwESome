package hash

type SHA1 struct{}

func (sha1 SHA1) Init() {

}
func (sha1 SHA1) Sum(b []byte) []byte {
	return b
}
