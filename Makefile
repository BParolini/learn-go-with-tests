.PHONY: tests bench

tests:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Done."

bench:
	@echo "Running benchmarks..."
	@go test -bench=. ./...
	@echo "Done."
