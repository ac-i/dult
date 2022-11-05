.PHONY: go_mod_tidy
go_mod_tidy:
	@go mod tidy

.PHONY: go_test_example
go_test_example:
	@go mod tidy
	@go test -timeout 30s -run ^ExampleQuickDULT$

.PHONY: go_test_quick_dult
go_test_quick_dult:
	@go mod tidy
	@go test -timeout 30s -run ^TestQuickDULT$

.PHONY: go_bench_quick_dult
go_bench_quick_dult:
	@go mod tidy
	@go test -benchmem -bench .

.PHONY: protoc_generate
protoc_generate:
	@protoc -I . \
		--go_out . \
		--go_opt paths=source_relative \
		\
 		--go-grpc_out . \
		--go-grpc_opt paths=source_relative \
		\
 		--grpc-gateway_out . \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		\
 		--openapiv2_out proto/speedpb/ \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt enums_as_ints=true \
		--openapiv2_opt allow_merge=true \
		\
		proto/*/*.proto
	@go mod tidy

.PHONY: apidocs_generate
apidocs_generate:
	@protoc -I . \
		--doc_out . \
		--doc_opt markdown,apidocs.md,source_relative \
		proto/*/*.proto
	@protoc -I . \
		--doc_out . \
		--doc_opt html,apidocs.html,source_relative \
		proto/*/*.proto
	@protoc -I . \
		--doc_out . \
		--doc_opt json,apidocs.json,source_relative \
		proto/*/*.proto
	@protoc -I . \
 		--openapiv2_out proto/speedpb/ \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt enums_as_ints=true \
		--openapiv2_opt allow_merge=true \
		proto/*/*.proto
	@go mod tidy

.PHONY: go_install_dependencies
go_install_dependencies:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

.PHONY: go_install_other
go_install_other:
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/99designs/gqlgen@latest
	@go install github.com/martinxsliu/protoc-gen-graphql@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest \
								github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@latest \
								github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@latest

