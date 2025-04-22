lint:
	$(GO_LINT)
	$(SQUAWK) --exclude=$(SQUAWK_RULES) ./migrations/*.sql
	$(PROTOLINT) lint $(PROTO_DIR)

test:
	$(GO_TEST) -coverprofile=$(COVERAGE_FILE_NAME) ./...
	$(GO_TOOL) cover -html=$(COVERAGE_FILE_NAME) -o $(COVERAGE_HTML_FILE_NAME)
