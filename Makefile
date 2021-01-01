# Makefile for mvnv
GIT_COMMIT   = github.com/shohi/mvnv/cmd/mvnv/version.gitCommit
GIT_REVISION = $(shell git rev-parse --short HEAD)
GOENV        = CGO_ENABLED=0 GO111MODULE=on

install:
	@# install mvnv
	@cd cmd/mvnv && $(GOENV) go install -ldflags "-X $(GIT_COMMIT)=$(GIT_REVISION) " .

	@# install mvn wrapper
	@cd cmd/mvn && $(GOENV) go install .

.phony: install
