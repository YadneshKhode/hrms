#In Go (Golang), a Makefile is a special file that contains a set of directives used to build and manage a project. Makefiles are commonly used in software development to automate the build process and execute various tasks, such as compiling code, running tests, and managing dependencies. The @ symbol in a Makefile is used to suppress the echoing of the command line to the standard output. It is commonly used at the beginning of an action line, which is the line that specifies the commands to be executed when the target is made.

build:
	@go build -o bin/go-ab-framework

run: build
	@./bin/go-ab-framework

test:
	@go test -v ./...

# Add a target to clean generated binaries
clean:
	@rm -rf bin

# Add a target to install dependencies
deps:
	@go get -u ./...

# Add a target to format code
fmt:
	@go fmt ./...

# Add a target to lint code (you need to install a linter, e.g., 'golangci-lint')
lint:
	@golangci-lint run

.PHONY: build run test clean deps fmt lint
