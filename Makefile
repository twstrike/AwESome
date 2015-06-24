test: deps
	golint ./...
	go test -v ./... -cover

deps:
	go get gopkg.in/check.v1
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/cover

