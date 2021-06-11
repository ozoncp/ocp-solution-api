.PHONY: build
build: .generate

PHONY: .generate
.generate:
		mkdir -p swagger
		for model in solution verdict ; do \
  			echo "generating $$model..." ; \
  			mkdir -p pkg/ocp-$$model-api ; \
			protoc --go_out=pkg/ocp-$$model-api --go_opt=paths=import \
				   --go-grpc_out=pkg/ocp-$$model-api --go-grpc_opt=paths=import \
				   --grpc-gateway_out=pkg/ocp-$$model-api \
				   --grpc-gateway_opt=logtostderr=true \
				   --grpc-gateway_opt=paths=import \
				   --swagger_out=allow_merge=true,merge_file_name=ocp-$$model-api:swagger \
				   api/ocp-$$model-api/ocp-$$model-api.proto ; \
			mv pkg/ocp-$$model-api/github.com/ozoncp/ocp-solution-api/pkg/ocp-$$model-api/* pkg/ocp-$$model-api/ ; \
			rm -rf pkg/ocp-$$model-api/github.com ; \
		done
