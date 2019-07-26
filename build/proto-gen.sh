export GO_PATH=/Users/amehrotra/workspace/Go
export PROTO_PATH=$GO_PATH/src/bitbucket-repository-management-service/dm/proto-files

protoc -I $GO_PATH/src --go_out=$GO_PATH/src $PROTO_PATH/domain/repository.proto
protoc -I $GO_PATH/src --go_out=plugins=grpc:$GO_PATH/src $PROTO_PATH/service/repository-service.proto

#TODO

#protoc -I $GO_PATH/src --go_out=$GO_PATH/src $GO_PATH/src/bitbucket-repository-management-service/dm/proto-files/domain/todo.proto
#protoc -I $GO_PATH/src --go_out=plugins=grpc:$GO_PATH/src $GO_PATH/src/bitbucket-repository-management-service/dm/proto-files/service/todo-service.proto

protoc -I $GO_PATH/src --go_out=plugins=grpc:$GO_PATH/src  --proto_path=third_party $PROTO_PATH/domain/todo.proto


protoc -I $GO_PATH/src --go_out=plugins=grpc:$GO_PATH/src --proto_path=third_party $PROTO_PATH/service/todo-service.proto

protoc  --proto_path=$PROTO_PATH/service:$PROTO_PATH/domain --proto_path=third_party --grpc-gateway_out=logtostderr=true:$GO_PATH/src $PROTO_PATH/service/todo-service.proto

protoc  --proto_path=$PROTO_PATH/service:$PROTO_PATH/domain  --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/service $PROTO_PATH/service/todo-service.proto
#protoc -I $GO_PATH/src --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/domain $PROTO_PATH/domain/todo.proto

#protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto
#protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api/v1 todo-service.proto
#protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger/v1 todo-service.proto
