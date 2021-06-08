package main

import (
	api "gitlab.com/insanitywholesale/lister/grpc"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"gitlab.com/insanitywholesale/lister/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
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
	go grpcServer.Serve(listener)

	restport := os.Getenv("LISTER_REST_PORT")
	if restport == "" {
		restport = "9392"
	}
	log.Fatal(rest.RunGateway(grpcport, restport))
}
