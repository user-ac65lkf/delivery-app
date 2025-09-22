test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

app_run:
	docker-compose -f docker-compose.yaml down -v
	docker-compose -f docker-compose.yaml up -d db
	sleep 15
	docker-compose -f docker-compose.yaml up -d init
	docker-compose -f docker-compose.yaml up -d migrate
	docker-compose -f docker-compose.yaml up -d uzum_delivery
#	docker-compose -f docker-compose.yaml up -d zookeeper
#	docker-compose -f docker-compose.yaml up -d kafka
#	docker exec -it uzum_delivery-kafka-1 kafka-topics --create --bootstrap-server localhost:29092 --topic my_topic --partitions 3 --replication-factor 1

app_run2:
	docker-compose -f docker-compose.yaml down -v
	docker-compose up -d
#Записать в Топик
#docker exec -it uzum_delivery-kafka-1 kafka-console-producer --broker-list localhost:29092 --topic my_topic

BIN:=$(CURDIR)/bin
install-deps:
	GOBIN=$(BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

gen: ## Генерация proto-файлов
		mkdir -p pkg/login_v1
		protoc 	--proto_path=api/login_v1 \
				--proto_path=proto \
				--go_out=pkg/login_v1 --go_opt=paths=source_relative \
				--plugin=protoc-gen-go=bin/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
				--go-grpc_out=pkg/login_v1 --go-grpc_opt=paths=source_relative \
				api/login_v1/login.proto