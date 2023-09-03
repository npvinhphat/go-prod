# Go project settings
GO := go
BINARY_NAME := go-prod
DOC_OUTPUT_DIR := docs/examples

# Build the project
build:
	$(GO) build -o $(BINARY_NAME) main.go

# Run the project
run:
	$(GO) run main.go

# Build the project and run it
run-build: build
	./$(BINARY_NAME)

# Generate examples
generate-examples:
	@mkdir -p $(DOC_OUTPUT_DIR)
	./$(BINARY_NAME) generate checklist --stacks=default,kubernetes,pagerduty,slack,wavefront --level=a > $(DOC_OUTPUT_DIR)/checklist.md
	./$(BINARY_NAME) generate guideline --stacks=default,kubernetes,pagerduty,slack,wavefront > $(DOC_OUTPUT_DIR)/guideline.md
	@echo "Examples generated in $(DOC_OUTPUT_DIR)/."
