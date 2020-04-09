export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

ATOMIX_RAFT_STORAGE_VERSION := latest

all: build

build: # @HELP build the source code
build: deps
	GOOS=linux GOARCH=amd64 go build -o build/_output/raft-storage ./cmd/raft-storage
	GOOS=linux GOARCH=amd64 go build -o build/_output/raft-storage-controller ./cmd/raft-storage-controller


deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

test: # @HELP run the unit tests and source code validation
test: build license_check linters
	go test github.com/atomix/raft-storage/...

coverage: # @HELP generate unit test coverage data
coverage: build linters license_check
	go test github.com/atomix/raft-storage/pkg/... -coverprofile=coverage.out.tmp -covermode=count
	@cat coverage.out.tmp | grep -v ".pb.go" > coverage.out

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

license_check: # @HELP examine and ensure license headers exist
	./build/licensing/boilerplate.py -v

proto: # @HELP build Protobuf/gRPC generated types
proto:
	docker run -it -v `pwd`:/go/src/github.com/atomix/raft-storage \
		-w /go/src/github.com/atomix/raft-storage \
		--entrypoint build/bin/compile_protos.sh \
		onosproject/protoc-go:stable

images: # @HELP build atomix storage and atomix storage controller Docker images
images: build
	docker build . -f build/raft-storage/Dockerfile -t atomix/raft-storage:${ATOMIX_RAFT_STORAGE_VERSION}
	docker build . -f build/raft-storage-controller/Dockerfile -t atomix/raft-storage-controller:${ATOMIX_RAFT_STORAGE_VERSION}

push: # @HELP push atomix-raft-node Docker image
	docker push atomix/raft-storage:${ATOMIX_RAFT_STORAGE_VERSION}
	docker push atomix/raft-storage-controller:${ATOMIX_RAFT_STORAGE_VERSION}
