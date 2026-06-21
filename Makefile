BINARY_NAME=redis-ui
MAIN_PATH=./main.go
VERSION=$(shell cat version)
GO=/usr/local/go/bin/go
DOCKER_IMAGE=redis-ui

.PHONY: deps build run dev clean kill docker-build docker-run

deps:
	$(GO) mod download
	$(GO) mod tidy

build:
	$(GO) build -o build/$(BINARY_NAME) $(MAIN_PATH)

run:
	$(GO) run $(MAIN_PATH)

dev:
	find . -name "*.go" | entr -r $(GO) run $(MAIN_PATH)

kill:
	-fuser -k 8080/tcp

docker-build:
	docker build -f ops/Dockerfile -t $(DOCKER_IMAGE):$(VERSION) -t $(DOCKER_IMAGE):latest .

docker-run:
	docker run --rm -p 8080:8080 $(DOCKER_IMAGE):latest

clean:
	rm -rf build/
