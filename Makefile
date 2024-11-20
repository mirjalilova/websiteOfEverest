CURRENT_DIR=$(shell pwd)
APP=everest
APP_CMD_DIR=./cmd

DB_URL="postgres://postgres:1111@localhost/everest?sslmode=disable"

run:
	go run cmd/main.go

init:
	# go mod init
	go mod tidy 
	go mod vendor


migrate_up:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/everest?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/everest?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/everest?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

insert_file:
	migrate create -ext sql -dir migrations -seq insert_table

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}