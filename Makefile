.PHONY: run-be
run-be:
	@go run main.go

.PHONY: run-be-dev
run-be-dev:
	@go run main.go

.PHONY: build
build:
	@make build-be build-fe -j

.PHONY: run
run:
	@./goblog

.PHONY: build-be
build-be:
	@go build -o goblog

.PHONY: build-fe
build-fe:
	@cd client; npm run build