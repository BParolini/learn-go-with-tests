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

.PHONY: tests bench cov vet
