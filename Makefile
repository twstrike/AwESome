test: deps
	go test -v ./...

deps:
	go get gopkg.in/check.v1