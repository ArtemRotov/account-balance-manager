.PHONY: run
run:
	go run ./cmd/app/main.go
 
.PHONY: build
build:
	go build -v -o app_service ./cmd/app/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: swag
swag: ### generate swagger docs
	swag init -g internal/app/app.go --parseInternal --parseDependency

.DEFAULT_GOAL := build