GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

# generate errors proto
.PHONY: errors
errors:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
		   --go-errors_out=paths=source_relative:./api \
	       $(API_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest

.PHONY: initProject
initProject:
# 初始化项目，将package改为project, 即 module github.com/medtrib/$(PROJECT)
	@echo "PROJECT=$(PROJECT)"
	@grep  -rl $(PROJECT) ./ | xargs sed -i 's/$(PROJECT)/github.com\/ZQCard\/$(PROJECT)/g'

.PHONY: initNewService
initNewService:
# 复制proto文件
	@mkdir -p ./api/$(ServiceLowerName)/v1/
	@cp ./api/admin/v1/admin.proto ./api/$(ServiceLowerName)/v1/$(ServiceLowerName).proto

# 删除旧文件
	@rm -rf ./api/admin
# 生成proto客户端文件
	@ kratos proto client ./api/$(ServiceLowerName)/v1/$(ServiceLowerName).proto
# 替换 Admin $(ServiceUpperName), admin 为$(ServiceLowerName)并迁移文件
	@grep  -rl Admin ./ | xargs sed -i 's/Admin/$(ServiceUpperName)/g'
	@grep  -rl admin ./ | xargs sed -i 's/admin/$(ServiceLowerName)/g'

# 迁移文件
	@ mv ./internal/service/admin.go ./internal/service/$(ServiceLowerName).go
	@ mv ./internal/biz/admin.go ./internal/biz/$(ServiceLowerName).go
	@ mv ./internal/data/admin.go ./internal/data/$(ServiceLowerName).go
	@ mv ./internal/domain/admin.go ./internal/domain/$(ServiceLowerName).go
# 拉取引用包
	go mod tidy
	@echo "project start success"

.PHONY: all
# generate all
all:
	make api
	make errors
	make generate

.DEFAULT_GOAL := all
