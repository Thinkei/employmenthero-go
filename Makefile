all: test vet lint

build:
	go build ./...

vet:
	go vet ./...

test:
	go test ./...

lint:
	staticcheck

codegen-format:
	go fmt ./...
