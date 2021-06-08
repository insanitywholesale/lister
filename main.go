package main

import (
	"fmt"
	api "gitlab.com/insanitywholesale/lister/grpc"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hey")
	grpcport := os.Getenv("LISTER_GRPC_PORT")
	if grpcport == "" {
		grpcport = "15200"
	}

	listener, err := net.Listen("tcp", ":"+grpcport)
	if err != nil {
		log.Fatalf("listen failed %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterListerServer(grpcServer, api.Server{})
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(listener))
}
