.PHONY: swagger build up down lb fmt test cover all rs

swagger:
	@$(CURDIR)/scripts/gen-swagger.sh

build:
	@docker compose build

up:
	@docker compose up -d

down:
	@docker compose down --remove-orphans

lb:
	@go build -o . ./...

fmt:
	@go fmt ./...

test:
	@go test ./...

cover:
	@go test -cover -coverprofile=c.out ./...
	@go tool cover -html=c.out -o coverage.html

all: build up

rs: down up

.PHONY: build-integ
build-integ:
	@docker build -t mikosurge/grace:latest -f ./build/grace/Dockerfile .
	# @docker build -t mikosurge/mockserver:latest -f ./build/mockserver/Dockerfile .
