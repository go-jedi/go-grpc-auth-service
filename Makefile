LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/air-verse/air@latest

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

update-packages:
	go get -u ./...

run-air:
	air --build.cmd "go build -o .bin/air cmd/app/main.go" --build.bin "./.bin/air --config testdata/config.yaml"

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.yaml

generate-proto:
	protoc --proto_path gen \
	--go_out=gen --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	-I . \
	proto/model/v1/user.proto proto/service/v1/auth.proto proto/service/v1/user.proto

mock-generate:
	rm -rf internal/service/mocks
	mockgen -source=internal/service/service.go \
	-destination=internal/service/mocks/mock_service.go

	rm -rf internal/repository/mocks
	mockgen -source=internal/repository/repository.go \
    	-destination=internal/repository/mocks/mock_repository.go

test-coverage:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out