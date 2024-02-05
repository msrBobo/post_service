DB_URL := "postgres://postgres:1234@localhost:5432/userdb?sslmode=disable"

CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	protoc --go_out=genproto --go-grpc_out=genproto protos/user.proto	
	
run:
	go run cmd/main.go

migrate-up:
    migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
    migrate -path migrations -database "$(DB_URL)" -verbose down

migrate_file:
    migrate create -ext sql -dir migrations/ -seq users

migrate-dirty:
	migrate -path ./migrations/ -database "postgresql://postgres:1234@localhost:5432/userdb?sslmode=disable" force 1