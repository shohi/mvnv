# Makefile for mvnv
GIT_COMMIT   = github.com/shohi/mvnv/cmd/mvnv/version.gitCommit
GIT_REVISION = $(shell git rev-parse --short HEAD)
GOENV        = CGO_ENABLED=0 GO111MODULE=on

MVNV_BINARY         = bin/mvnv
MVN_WRAPPER_BINARY  = bin/mvn

.phony: install
install:
	@# install mvnv
	@cd cmd/mvnv && $(GOENV) go install -ldflags "-X $(GIT_COMMIT)=$(GIT_REVISION) " .

	@# install mvn wrapper
	@cd cmd/mvn && $(GOENV) go install .

.phony: build
build:
	@# build mvnv
	$(GOENV) go build \
		-ldflags "-X $(GIT_COMMIT)=$(GIT_REVISION) " \
		-o $(MVNV_BINARY) \
		./cmd/mvnv \

	@# build mvn wrapper
	$(GOENV) go build \
		-o $(MVN_WRAPPER_BINARY) \
		./cmd/mvn
