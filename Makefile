SHELL = /bin/bash -o pipefail

STATICCHECK := $(GOPATH)/bin/staticcheck
version ?= minor

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

.PHONY: release
release: race-test
	go run github.com/kevinburke/bump_version@latest --tag-prefix=v $(version) restclient/version.go

ci: test-deps lint race-test
