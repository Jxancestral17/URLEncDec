build:
	@go build -o bin/urlencdec


run: build 
	@./bin/urlencdec

test:
	@go test ./... -v