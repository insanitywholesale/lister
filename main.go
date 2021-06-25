package main

import (
	_ "embed"
	api "gitlab.com/insanitywholesale/lister/grpc"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"gitlab.com/insanitywholesale/lister/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var (
	//go:embed openapiv2/v1/lister.swagger.json
	openapiDocs []byte

	commitHash string
	commitDate string
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
	log.Println("grpc started on port", grpcport)
	go grpcServer.Serve(listener)

	restport := os.Getenv("LISTER_REST_PORT")
	if restport == "" {
		restport = "9392"
	}

	rest.SaveVars(openapiDocs, commitHash, commitDate)

	log.Println("rest starting on port", restport)
	log.Fatal(rest.RunGateway(grpcport, restport))
}
