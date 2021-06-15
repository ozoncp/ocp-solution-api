MODELS = solution verdict
MOCK_PACKAGES = flusher saver
OS=$(shell uname | tr A-Z a-z)

.PHONY: build
build: .generate-mocks .deps .vendor-proto .generate-api .build lint

.PHONY: .build
.build:
	CGO_ENABLED=0 GOOS=$(OS) go build -o bin/ocp-solution-api cmd/ocp-solution-api/main.go || \
	CGO_ENABLED=0 GOOS=windows go build -o bin/ocp-solution-api cmd/ocp-solution-api/main.go
	go mod tidy

.PHONY: .generate-api
.generate-api:
	mkdir -p swagger
	for model in $(MODELS) ; do \
		echo "generating $$model ..." ; \
		mkdir -p pkg/ocp-$$model-api ; \
		protoc -I vendor.protogen \
		       --go_out=pkg/ocp-$$model-api --go_opt=paths=import \
			   --go-grpc_out=pkg/ocp-$$model-api --go-grpc_opt=paths=import \
			   --grpc-gateway_out=pkg/ocp-$$model-api \
			   --grpc-gateway_opt=logtostderr=true \
			   --grpc-gateway_opt=paths=import \
			   --validate_out lang=go:pkg/ocp-$$model-api \
			   --swagger_out=allow_merge=true,merge_file_name=ocp-$$model-api:swagger \
			   api/ocp-$$model-api/ocp-$$model-api.proto ; \
		mv pkg/ocp-$$model-api/github.com/ozoncp/ocp-solution-api/pkg/ocp-$$model-api/* pkg/ocp-$$model-api/ ; \
		rm -rf pkg/ocp-$$model-api/github.com ; \
	done

.PHONY: .vendor-proto
.vendor-proto: vendor.protogen/google vendor.protogen/github.com/envoyproxy
	mkdir -p vendor.protogen
	for model in $(MODELS) ; do \
  		mkdir -p vendor.protogen/api/ocp-$$model-api ; \
  		cp api/ocp-$$model-api/ocp-$$model-api.proto vendor.protogen/api/ocp-$$model-api ; \
	done

vendor.protogen/google:
	git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis
	mkdir -p vendor.protogen/google
	mv vendor.protogen/googleapis/google/api vendor.protogen/google
	rm -rf vendor.protogen/googleapis

vendor.protogen/github.com/envoyproxy:
	mkdir -p vendor.protogen/github.com/envoyproxy
	git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate

.PHONY: .generate-mocks
.generate-mocks:
	for package in $(MOCK_PACKAGES) ; do \
		cd internal/$$package ; \
		ginkgo bootstrap ; \
		rm -f $$package_test.go ; \
		ginkgo generate $$package ; \
		cd - ; \
	done
	go generate ./...

.PHONY: .deps
.deps:
	ls go.mod || go mod init
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/envoyproxy/protoc-gen-validate

.PHONY: postgres
postgres:
	pg_ctl -D /usr/local/var/postgres start

.PHONY: lint
lint:
	golangci-lint run
