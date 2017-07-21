BUMP_VERSION := $(shell command -v bump_version)
STATICCHECK := $(shell command -v megacheck)

vet:
ifndef STATICCHECK
	go get -u honnef.co/go/tools/cmd/megacheck
endif
	go vet ./...
	megacheck ./...

test: vet
	bazel test //...

race-test: vet
	go test -race ./...

release: race-test
ifndef BUMP_VERSION
	go get github.com/Shyp/bump_version
endif
	bump_version minor client.go
