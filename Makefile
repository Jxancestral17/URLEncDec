build:
	@go build -o bin/JXSS


run: build 
	@./bin/JXSS

test:
	@go test ./... -v