build:
	@go build -o ./bin/bid
	
run: build
	@./bin/bid
	
test:
	@go test ./...