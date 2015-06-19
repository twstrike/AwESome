test: deps
	golint ./...
	go test -v ./...

deps:
	go get gopkg.in/check.v1
	go get github.com/golang/lint/golint

