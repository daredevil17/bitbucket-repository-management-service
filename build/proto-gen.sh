export GO_PATH=/Users/amehrotra/workspace/Go
protoc -I $GO_PATH/src --go_out=$GO_PATH/src $GO_PATH/src/bitbucket-repository-management-service/dm/proto-files/domain/repository.proto
protoc -I $GO_PATH/src --go_out=plugins=grpc:$GO_PATH/src $GO_PATH/src/bitbucket-repository-management-service/dm/proto-files/service/repository-service.proto
