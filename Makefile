gen:
	go generate ./...

test:
	go test -cover ./...
	golangci-lint run --print-issued-lines=false --out-format=colored-line-number --issues-exit-code=0 ./...

test-cover:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

build:
	go build cmd/bot/main.go