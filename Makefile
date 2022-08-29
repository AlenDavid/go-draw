mod:
	go mod tidy

build: mod
	@go build cmd/go-draw/go-draw.go

run: mod
	go run cmd/go-draw/go-draw.go
