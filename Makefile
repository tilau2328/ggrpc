projects := $(notdir $(foreach dir,./proto,$(wildcard $(dir)/*)))
js_rpc_out := --go_out=plugins=grpc:.
go_rpc_out := --go_out=plugins=grpc:.
go_out := --go_out=plugins=grpc:.
pb_in := -I proto/

.PHONY: help
help:
	@echo "Usage: make <TARGET>"
	@echo ""
	@echo "Available targets are:"
	@echo ""
	@echo "    grpc-go              Generate Go files from proto"
	@echo "    generate-js              Generate JavaScript files from proto"

.PHONY: grpc
grpc: grpc-go grpc-js

.PHONY: grpc-js
grpc-js:
ifeq ($strip($(prj)),)
	foreach prj, $(projects), protoc $(pb_in) $(js_out) $(js_rpc_out) \
		proto/$(prj)/command/*.proto \
		proto/$(prj)/domain/*.proto \
		proto/$(prj)/query/*.proto \
		proto/$(prj)/event/*.proto
else
	protoc $(pb_in) $(js_out) $(js_rpc_out) \
		proto/$(prj)/command/*.proto \
		proto/$(prj)/domain/*.proto \
		proto/$(prj)/query/*.proto \
		proto/$(prj)/event/*.proto
endif

.PHONY: grpc-go
grpc-go:
ifeq ($strip($(prj)),)
	foreach prj, $(projects), \
		protoc $(pb_in) $(go_out) $(go_rpc_out) \
			proto/$(prj)/command/*.proto \
			proto/$(prj)/domain/*.proto \
			proto/$(prj)/query/*.proto \
			proto/$(prj)/event/*.proto
else
	protoc $(pb_in) $(go_out) $(go_rpc_out) \
		proto/$(prj)/command/*.proto \
		proto/$(prj)/domain/*.proto \
		proto/$(prj)/query/*.proto \
		proto/$(prj)/event/*.proto
endif
