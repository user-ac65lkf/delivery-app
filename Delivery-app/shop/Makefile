test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

BIN:=$(CURDIR)/bin

install:
	GOBIN=$(BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2


gen_shop: ## Генерация proto-файлов
		mkdir -p pkg/shopV1
		protoc 	--proto_path=api/shopV1 \
				--proto_path=proto \
				--go_out=pkg/shopV1 --go_opt=paths=source_relative \
				--plugin=protoc-gen-go=bin/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
				--go-grpc_out=pkg/shopV1 --go-grpc_opt=paths=source_relative \
				--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
				--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
				--grpc-gateway_out=pkg/shopV1 --grpc-gateway_opt=paths=source_relative \
				--openapiv2_out=allow_merge=true,merge_file_name=api_shopV1:docs \
                --plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
				api/shopV1/shop.proto

gen_login: ## Генерация auth proto-файлов
		mkdir -p pkg/loginV1
		protoc 	--proto_path=api/loginV1 \
				--proto_path=proto \
				--go_out=pkg/loginV1 --go_opt=paths=source_relative \
				--plugin=protoc-gen-go=bin/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
				--go-grpc_out=pkg/loginV1 --go-grpc_opt=paths=source_relative \
				api/loginV1/login.proto

mig_up:
	migrate -database postgres://postgres:postgres@localhost:5432/service?sslmode=disable -path migrations up

mig_down:
	migrate -database postgres://postgres:postgres@localhost:5432/service?sslmode=disable -path migrations down

create_topic:
	docker exec -it shopc-kafka-1 kafka-topics --create --bootstrap-server localhost:29092 --topic my_topic --partitions 3 --replication-factor 1

#write message to topic
write_message:
	docker exec -it shopc-kafka-1 kafka-console-producer --broker-list localhost:29092 --topic my_topic
