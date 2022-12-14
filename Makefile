all: test vet lint

build:
	go build ./...

vet:
	go vet ./...

test:
	go test ./...

lint:
	staticcheck

coverage:
	go test -covermode=count -coverprofile=combined.coverprofile ./...

coveralls:
	go install github.com/mattn/goveralls@latest && $(HOME)/go/bin/goveralls -service=github -coverprofile=combined.coverprofile

codegen-format:
	go fmt ./...
