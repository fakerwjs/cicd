.PHONY: build run test docker

build:
	go build -o app ./cmd/app

run:
	go run ./cmd/app

test:
	go test ./... -v

docker:
	docker build -t go-cicd-template .
