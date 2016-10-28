BUMP_VERSION := $(shell command -v bump_version)

vet:
	go vet ./...

test: vet
	go test ./...

race-test: vet
	go test -race ./...

release: race-test
ifndef BUMP_VERSION
	go get github.com/Shyp/bump_version
endif
	bump_version minor client.go
