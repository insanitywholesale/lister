package main

import (
	_ "embed"
	apiv1 "gitlab.com/insanitywholesale/lister/grpc/v1"
	pbv1 "gitlab.com/insanitywholesale/lister/proto/v1"
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

	grpcport string
	restport string

	serviceName string = "lister"
)

func getServiceName() string {
	return serviceName
}

func setupPorts() {
	grpcport = os.Getenv("LISTER_GRPC_PORT")
	if grpcport == "" {
		grpcport = "15200"
	}
	restport = os.Getenv("LISTER_REST_PORT")
	if restport == "" {
		restport = "9392"
	}
}

func startGRPC() {
	listener, err := net.Listen("tcp", ":"+grpcport)
	if err != nil {
		log.Fatalf("listen failed %v", err)
	}

	grpcServer := grpc.NewServer()
	pbv1.RegisterListerServer(grpcServer, apiv1.Server{})
	reflection.Register(grpcServer)
	log.Println("grpc started on port", grpcport)
	log.Fatal(grpcServer.Serve(listener))
}

func startHTTP() {
	rest.SaveVars(openapiDocs, commitHash, commitDate)

	log.Println("rest starting on port", restport)
	log.Fatal(rest.RunGateway(grpcport, restport))
}

func main() {
	setupPorts()
	go startGRPC()
	defer startHTTP()
}
