SHELL = /bin/bash

BUMP_VERSION := $(shell command -v bump_version)
MEGACHECK := $(shell command -v megacheck)

vet:
ifndef MEGACHECK
	go get -u honnef.co/go/tools/cmd/megacheck
endif
	go vet ./...
	megacheck ./...

test: vet
	bazel test \
		--remote_rest_cache=https://remote.rest.stackmachine.com/cache \
		--spawn_strategy=remote \
		--strategy=Closure=remote \
		--strategy=Javac=remote \
		--test_output=errors //...

race-test: vet
	bazel test \
		--remote_rest_cache=https://remote.rest.stackmachine.com/cache \
		--spawn_strategy=remote \
		--strategy=Closure=remote \
		--strategy=Javac=remote \
		--test_output=errors --features=race //...

ci:
	bazel --host_jvm_args=-Dbazel.DigestFunction=SHA1 test \
		--noshow_progress --noshow_loading_progress --test_output=errors \
		--features=race //...

release: race-test
ifndef BUMP_VERSION
	go get github.com/Shyp/bump_version
endif
	bump_version minor client.go
