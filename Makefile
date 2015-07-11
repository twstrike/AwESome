default: deps lint test

bench:
	go test -v ./... -check.b

lint:
	golint ./...

test:
	go test -v ./... -cover

deps:
	./deps.sh
