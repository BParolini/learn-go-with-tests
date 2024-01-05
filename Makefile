.PHONY: tests

tests:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Done."
