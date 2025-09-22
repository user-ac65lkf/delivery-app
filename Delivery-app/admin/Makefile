test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

BIN:=$(CURDIR)/bin
install-deps:
	GOBIN=$(BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	GOBIN=$(BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

gen: ## Генерация proto-файлов
		mkdir -p pkg/login_v1
		protoc 	--proto_path=api/login_v1 \
				--proto_path=proto \
				--go_out=pkg/login_v1 --go_opt=paths=source_relative \
				--plugin=protoc-gen-go=bin/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
				--go-grpc_out=pkg/login_v1 --go-grpc_opt=paths=source_relative \
				api/login_v1/login.proto
