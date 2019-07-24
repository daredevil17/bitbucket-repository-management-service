package main

import (
	cmd "bitbucket-repository-management-service/pkg/cmd/server"
	"fmt"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
