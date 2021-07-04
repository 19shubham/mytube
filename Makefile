.PHONY: build fmt run vet

MGSP_GENERATOR = $(GOBIN)/msgp
ROOT_DIR=${PWD}

GOPATH = ${ROOT_DIR}/vendor:${ROOT_DIR}

default: fmt clean build

build: vet
	go build -v -o ./bin/mytube

dev_rebuild: clean
	go build -v -o ./bin/mytube

vet:
	go vet ./src/...

clean:
	rm -rf `find ./vendor/src -type d -name .git` \
	&& rm -rf `find ./vendor/src -type d -name .hg` \
	&& rm -rf `find ./vendor/src -type d -name .bzr` \
	&& rm -rf `find ./vendor/src -type d -name .svn` \
	&& rm -rf `find . -type f -name *.pyc`
	rm -rf ./bin/mytube
	rm -rf ./pkg/*


# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@echo ${PATH}
	go fmt ./src/...
