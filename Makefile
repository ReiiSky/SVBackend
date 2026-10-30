# Variables
GO_CMD=go
MAIN_FILE=cmd/rest/*.go
PORT=3000
CONTAINER_WORKDIR=/app

# Default target
all: run

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o SVBackend $(MAIN_FILE)

# Run the main file
run:
	$(GO_CMD) run $(MAIN_FILE)

run-container:
	docker build -t sv-backend .
	docker run -d -p ${PORT}:${PORT} -v $(shell pwd)/.env:${CONTAINER_WORKDIR}/.env --name sv-backend-server --rm sv-backend

new-migrate:
	touch migrations/$(shell date +%s)_$(action).up.sql
	touch migrations/$(shell date +%s)_$(action).down.sql

deploy-migrate:
	migrate -source file://./migrations/  -database '$(url)' up
