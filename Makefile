test:
	@echo "Running the full test..."
	@go test -v -cover -race ./... -coverprofile cover.out