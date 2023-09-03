# Go project settings
GO := go
BINARY_NAME := go-prod
DOC_OUTPUT_DIR := docs/examples

# Build the project
build:
	$(GO) build -o $(BINARY_NAME) main.go

# Run the project
run: build
	$(GO) run main.go

# Generate examples
generate-examples: build
	@mkdir -p $(DOC_OUTPUT_DIR)
	./$(BINARY_NAME) generate checklist --level=a > $(DOC_OUTPUT_DIR)/default_checklist.md
	./$(BINARY_NAME) generate guideline > $(DOC_OUTPUT_DIR)/default_guideline.md
	./$(BINARY_NAME) generate checklist --stacks=default,kubernetes,pagerduty,slack,wavefront --level=c > $(DOC_OUTPUT_DIR)/detailed_checklist.md
	./$(BINARY_NAME) generate guideline --stacks=default,kubernetes,pagerduty,slack,wavefront > $(DOC_OUTPUT_DIR)/detailed_guideline.md
	@echo "Examples generated in $(DOC_OUTPUT_DIR)/."
