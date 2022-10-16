BIN_NAME=poogie

VERSION?=development

all: build

install: go-get

run:
	go run main.go

build:
	go mod tidy && \
	GOOS=linux GOARCH=amd64 go build -o build/$(BIN_NAME) -a -v -tags musl

up:
	docker-compose up -d

rebuild:
	docker-compose up -d --build

clean:
	rm -rf build/$(BIN_NAME)

go-get:
	go mod download

coverage:
	go test -v ./... -coverprofile cover.out
	go tool cover -func cover.out

bump:
	version-bump patch