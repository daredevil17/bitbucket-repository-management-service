# bitbucket-repository-management-service
#server start
#go run todo-main.go -grpc-port=9090 -http-port=8080 -db-host=127.0.0.1:3306 -db-user=root -db-password=<pass> -db-schema=testdb

#client rest test
#cd bitbucket-repository-management-service/cmd/client-rest
#go run main.go -server=http://localhost:8080

#client grpc test
#cd bitbucket-repository-management-service/cmd/grpc/client/
#go run main-todo.go -server=localhost:9090