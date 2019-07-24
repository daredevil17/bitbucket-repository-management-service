package cmd

import (
	"bitbucket-repository-management-service/dm/grpc/impl"
	"bitbucket-repository-management-service/pkg/protocol/grpc"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	log.Println("Inside RunServer...")
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	fmt.Println(cfg.DatastoreDBUser)
	fmt.Println(cfg.DatastoreDBPassword)
	fmt.Println(cfg.DatastoreDBHost)
	fmt.Println(cfg.DatastoreDBSchema)

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)

	log.Println(dsn)

	db, err := sql.Open("mysql", dsn)
	//db, err := sql.Open("mysql", "root:test123@tcp(172.17.0.1:3306)/todo_server")

	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	/*
		stmtIns, err := db.Query("insert into test values('John', 'D.', 'Doe')")
		if err != nil {
			panic(err)
		}
		// Close the statement when we leave main() / the program terminates
		defer stmtIns.Close()
	*/
	defer db.Close()

	v1API := impl.NewToDoServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
