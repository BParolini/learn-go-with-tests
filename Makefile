.PHONY: tests bench cov

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
