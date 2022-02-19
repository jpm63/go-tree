all: cover build

test:
	go test -v ./...

cover:
	go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o=cover.html

build:
	go build
