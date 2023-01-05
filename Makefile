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
	go install github.com/mattn/goveralls@latest && $(HOME)/go/bin/goveralls -service=circle-ci -coverprofile=combined.coverprofile -repotoken=$(COVERALLS_TOKEN)

MAJOR_VERSION := $(shell echo $(VERSION) | sed 's/\..*//')
update-version:
	@echo "${VERSION}" > VERSION
	@perl -pi -e 's|const clientVersion = "[.\d\-\w]+"|const clientVersion = "$(VERSION)"|' employmenthero.go

codegen-format:
	go fmt ./...

.PHONY: codegen-format update-version

