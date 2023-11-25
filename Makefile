BINARY_NAME = task
.DEFAULT_GOAL := run

.PHONY: run build clean

run: build
	@./build/$(BINARY_NAME)

build:
	@go build -o ./build/$(BINARY_NAME)

clean:
	@go clean
	@rm -rf ./build
	@rm ./task.db
	@rm task.db.lock