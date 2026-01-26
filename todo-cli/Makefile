.PHONY: build install test clean run

BINARY_NAME=todo

build:
	@echo "Building..."
	go build -o bin/$(BINARY_NAME) ./cmd/todo

install:
	@echo "Installing..."
	go install ./cmd/todo

test:
	@echo "Testing..."
	go test ./...

clean:
	@echo "Cleaning..."
	rm -rf bin/ data/todos.json

run: build
	@echo "Running..."
	./bin/$(BINARY_NAME) $(ARGS)

# Example usage: make add ARGS="add \"Buy milk\""
add: build
	./bin/$(BINARY_NAME) add $(ARGS)

list: build
	./bin/$(BINARY_NAME) list