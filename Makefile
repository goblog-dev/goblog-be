.PHONY: run
run:
	@./goblog

.PHONY: run-dev
run-be-dev:
	@air

.PHONY: build
build:
	@go build -o goblog
