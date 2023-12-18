.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go
test:
	go test ./...
run: build
	./.bin/bot

