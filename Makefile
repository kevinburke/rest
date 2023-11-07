SHELL = /bin/bash -o pipefail

BUMP_VERSION := $(GOPATH)/bin/bump_version
STATICCHECK := $(GOPATH)/bin/staticcheck

deps:
	go get -u ./...

test-deps:
	go get -t -v ./...

test: lint
	go test ./...

lint: | $(STATICCHECK)
	go vet ./...
	$(STATICCHECK) ./...

$(STATICCHECK):
	go install honnef.co/go/tools/cmd/staticcheck@latest

race-test: lint
	go test -race ./...

$(BUMP_VERSION):
	go install github.com/kevinburke/bump_version

release: race-test | $(BUMP_VERSION)
	$(BUMP_VERSION) minor restclient/client.go

ci: test-deps lint race-test
