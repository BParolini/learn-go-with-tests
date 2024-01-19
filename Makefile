tests:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Done."

bench:
	@echo "Running benchmarks..."
	@go test -bench=. ./...
	@echo "Done."

cov:
	@echo "Running coverage..."
	@go test -v -cover ./...
	@echo "Done."

vet:
	@echo "Running vet..."
	@go vet -v ./...
	@echo "Done."

build:
	@echo "Building project..."
	@go build -o bin/learn-go-with-tests
	@echo "Done."

run_clockface: build
	@echo "Running Clockface..."
	@bin/learn-go-with-tests clockface
	@echo
	@echo "Done."

.PHONY: tests bench cov vet build \
	run_clockface
