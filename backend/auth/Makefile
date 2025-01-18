PROTO_DIR := proto

# Find all .proto files recursively in PROTO_DIR
PROTO_FILES := $(shell find $(PROTO_DIR) -type f -name "*.proto")

# Define generated Go files for each proto file
GO_FILES := $(PROTO_FILES:.proto=.pb.go)
GRPC_GO_FILES := $(PROTO_FILES:.proto=_grpc.pb.go)

# Generate all Go files for proto
.PHONY: generate/proto
generate/proto: $(GO_FILES) $(GRPC_GO_FILES)

%.pb.go: %.proto
	@echo "Generating Go code for $<..."
	protoc -I $(dir $<) $< \
		--go_out=$(dir $<) --go_opt=paths=source_relative

%_grpc.pb.go: %.proto
	@echo "Generating gRPC Go code for $<..."
	protoc -I $(dir $<) $< \
		--go-grpc_out=$(dir $<) --go-grpc_opt=paths=source_relative

.PHONY: clean
clean:
	@echo "Cleaning generated Go files..."
	rm -f $(GO_FILES) $(GRPC_GO_FILES)