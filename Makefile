.PHONY: run build build-run run-locally

build:
	@docker build -t gustavo-villar/go-stress-test:v1 .

run:
	@docker run gustavo-villar/go-stress-test:v1 --url=https://google.com --requests=10 --concurrency=2

build-run: build run

run-local:
	@go run cmd/main.go --url=https://google.com --requests=100 --concurrency=2
