PROJECT=go-web-boilerplate

build:	
	go build -v -o $(PROJECT) ./cmd/web

run: build
	./$(PROJECT) api

migrate: build
	./$(PROJECT) migrate

fmt:
	go fmt ./...

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

test_coverage_html: test_coverage
	go tool cover -html=coverage.out